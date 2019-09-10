package vouched

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/vincent-petithory/dataurl"
)

const defaultContentType = "application/json"

// VouchedAPIClient is a generic base class for calling the vouched API
type VouchedAPIClient struct {
	Host     string
	Path     string
	Scheme   string
	Token    *string
	Username *string
	Password *string
}

// NewVouchedAPIClient initializes an VouchedAPIClient using the environment-configured API
// user and token to construct an HTTP basic authorization header for access to the Vouched API.
func NewVouchedAPIClient() (*VouchedAPIClient, error) {
	apiURL, err := url.Parse(vouchedAPIBaseURL)
	if err != nil {
		log.Warningf("Failed to parse vouched API base url; %s", err.Error())
		return nil, err
	}

	return &VouchedAPIClient{
		Host:   apiURL.Host,
		Scheme: apiURL.Scheme,
		Path:   "/graphql",
	}, nil
}

func (v *VouchedAPIClient) buildGraphQLPayload(request string, params map[string]interface{}) []byte {
	inputJSON, _ := json.Marshal(map[string]interface{}{
		"query":     request,
		"variables": params,
	})
	return inputJSON
}

func (v *VouchedAPIClient) sendRequest(method, uri, contentType, graphqlRequest string, params map[string]interface{}, response interface{}) (status int, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: time.Second * 30,
	}

	urlString := v.buildURL(uri)
	mthd := strings.ToUpper(method)
	reqURL, err := url.Parse(urlString)
	if err != nil {
		log.Warningf("Failed to parse URL for vouched API (%s %s) invocation; %s", method, urlString, err.Error())
		return -1, err
	}

	payload := v.buildGraphQLPayload(graphqlRequest, params)
	log.Debugf("Attempting graphQL request:\n%s", string(payload))

	if mthd == "GET" && params != nil {
		_params := map[string]interface{}{}
		json.Unmarshal(payload, &_params) // HACK
		q := reqURL.Query()
		for name := range _params {
			if val, valOk := params[name].(string); valOk {
				q.Set(name, val)
			}
		}
		reqURL.RawQuery = q.Encode()
	}

	headers := map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Accept":          {"application/json"},
	}
	if v.Username != nil && v.Password != nil {
		headers["Authorization"] = []string{buildBasicAuthorizationHeader(*v.Username, *v.Password)}
	} else if v.Token != nil {
		headers["Authorization"] = []string{fmt.Sprintf("Bearer %s", *v.Token)}
	}

	if vouchedAPIToken != "" {
		headers["X-Api-Key"] = []string{vouchedAPIToken}
	}

	var req *http.Request

	if mthd == "POST" || mthd == "PUT" {
		if contentType == "application/json" {
			// payload, err = json.Marshal(params)
			// FIXME-- this branch is a no-op due to vouched's graphql backend
			if err != nil {
				log.Warningf("Failed to marshal JSON payload for vouched API (%s %s) invocation; %s", method, urlString, err.Error())
				return -1, err
			}
		} else if contentType == "application/x-www-form-urlencoded" {
			urlEncodedForm := url.Values{}
			for key, val := range params {
				if valStr, valOk := val.(string); valOk {
					urlEncodedForm.Add(key, valStr)
				} else {
					log.Warningf("Failed to marshal application/x-www-form-urlencoded parameter: %s; value was non-string", key)
				}
			}
			payload = []byte(urlEncodedForm.Encode())
		} else if contentType == "multipart/form-data" {
			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			for key, val := range params {
				if valStr, valStrOk := val.(string); valStrOk {
					dURL, err := dataurl.DecodeString(valStr)
					if err == nil {
						log.Debugf("Parsed data url parameter: %s", key)
						part, err := writer.CreateFormFile(key, key)
						if err != nil {
							return 0, err
						}
						part.Write(dURL.Data)
					} else {
						_ = writer.WriteField(key, valStr)
					}
				} else {
					log.Warningf("Skipping non-string value when constructing multipart/form-data request: %s", key)
				}
			}
			err = writer.Close()
			if err != nil {
				return 0, err
			}

			payload = []byte(body.Bytes())
		}

		req, _ = http.NewRequest(method, urlString, bytes.NewReader(payload))
		headers["Content-Type"] = []string{contentType}
	} else {
		req = &http.Request{
			URL:    reqURL,
			Method: mthd,
		}
	}

	req.Header = headers

	resp, err := client.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Warningf("Failed to invoke vouched API (%s %s); %s", method, urlString, err.Error())
		return 0, err
	}

	log.Debugf("Received %v response for vouched API (%s %s) invocation", resp.StatusCode, method, urlString)

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	err = json.Unmarshal(buf.Bytes(), response)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("Failed to unmarshal vouched API (%s %s) response: %s; %s", method, urlString, buf.Bytes(), err.Error())
	}

	if resp.StatusCode < 400 {
		log.Debugf("Invocation of vouched graphQL API succeeded (%v-byte response)", buf.Len())
	} else {
		log.Warningf("Invocation of vouched graphQL API failed (%v-byte response)", buf.Len())
		return resp.StatusCode, fmt.Errorf("Vouched graphQL API invocation failed (%d)", resp.StatusCode)
	}
	return resp.StatusCode, nil
}

// Get constructs and synchronously sends an API GET request
func (v *VouchedAPIClient) Get(graphqlRequest string, params map[string]interface{}, response interface{}) (status int, err error) {
	return v.sendRequest("GET", "", defaultContentType, graphqlRequest, params, response)
}

// Post constructs and synchronously sends an API POST request
func (v *VouchedAPIClient) Post(graphqlRequest string, params map[string]interface{}, response interface{}) (status int, err error) {
	return v.sendRequest("POST", "", defaultContentType, graphqlRequest, params, response)
}

// PostWWWFormURLEncoded constructs and synchronously sends an API POST request using
func (v *VouchedAPIClient) PostWWWFormURLEncoded(graphqlRequest string, params map[string]interface{}, response interface{}) (status int, err error) {
	return v.sendRequest("POST", "", "application/x-www-form-urlencoded", graphqlRequest, params, response)
}

// PostMultipartFormData constructs and synchronously sends an API POST request using multipart/form-data as the content-type
func (v *VouchedAPIClient) PostMultipartFormData(graphqlRequest string, params map[string]interface{}, response interface{}) (status int, err error) {
	return v.sendRequest("POST", "", "multipart/form-data", graphqlRequest, params, response)
}

func (v *VouchedAPIClient) buildURL(uri string) string {
	return fmt.Sprintf("%s://%s%s/%s", v.Scheme, v.Host, v.Path, uri)
}

func buildBasicAuthorizationHeader(username, password string) string {
	auth := fmt.Sprintf("%s:%s", username, password)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth)))
}
