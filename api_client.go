package vouched

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	log.Debugf("Attempting %d-byte vouched graphQL API request", len(payload))

	if mthd == "GET" && params != nil {
		_params := map[string]interface{}{}
		json.Unmarshal(payload, &_params) // HACK
		q := reqURL.Query()
		for name := range _params {
			if val, valOk := _params[name].(string); valOk {
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
		log.Debugf("%s", string(buf.Bytes()))
	} else {
		log.Warningf("Invocation of vouched graphQL API failed (%v-byte response follows)\n%s", buf.Len(), string(buf.Bytes()))
		return resp.StatusCode, fmt.Errorf("Vouched graphQL API invocation failed (%d)", resp.StatusCode)
	}
	return resp.StatusCode, nil
}

// Post constructs and synchronously sends an API POST request
func (v *VouchedAPIClient) Post(graphqlRequest string, params map[string]interface{}, response interface{}) (status int, err error) {
	return v.sendRequest("POST", "", defaultContentType, graphqlRequest, params, response)
}

func (v *VouchedAPIClient) buildURL(uri string) string {
	return fmt.Sprintf("%s://%s%s/%s", v.Scheme, v.Host, v.Path, uri)
}

func buildBasicAuthorizationHeader(username, password string) string {
	auth := fmt.Sprintf("%s:%s", username, password)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth)))
}
