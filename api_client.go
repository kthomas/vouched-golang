package vouched

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/machinebox/graphql"
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
		Host:     apiURL.Host,
		Scheme:   apiURL.Scheme,
		Path:     "",
		Username: stringOrNil(vouchedAPIUser),
		Password: stringOrNil(vouchedAPIToken),
	}, nil
}

func (v *VouchedAPIClient) sendRequest(method, urlString, contentType string, params map[string]interface{}, response interface{}) (status int, err error) {
	client := graphql.NewClient(urlString)

	// TODO: params to graphql request
	graphqlRequest := ""

	req := graphql.NewRequest(graphqlRequest)
	req.Header.Set("X-Api-Key", vouchedAPIToken)

	if response == nil {
		response = &KYCApplicationIDVerificationResult{}
	}

	err = client.Run(context.TODO(), req, response)
	if err != nil {
		log.Warningf("Failed to execute graphql request; %s", err.Error())
		return -1, err
	}

	log.Debugf("Invocation of vouched graphql API succeeded; response: %s", response)
	return 0, nil // FIXME: return relevant status
}

// Get constructs and synchronously sends an API GET request
func (v *VouchedAPIClient) Get(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := v.buildURL(uri)
	return v.sendRequest("GET", url, defaultContentType, params, response)
}

// Post constructs and synchronously sends an API POST request
func (v *VouchedAPIClient) Post(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := v.buildURL(uri)
	return v.sendRequest("POST", url, defaultContentType, params, response)
}

// PostWWWFormURLEncoded constructs and synchronously sends an API POST request using
func (v *VouchedAPIClient) PostWWWFormURLEncoded(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := v.buildURL(uri)
	return v.sendRequest("POST", url, "application/x-www-form-urlencoded", params, response)
}

// PostMultipartFormData constructs and synchronously sends an API POST request using multipart/form-data as the content-type
func (v *VouchedAPIClient) PostMultipartFormData(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := v.buildURL(uri)
	return v.sendRequest("POST", url, "multipart/form-data", params, response)
}

// Put constructs and synchronously sends an API PUT request
func (v *VouchedAPIClient) Put(uri string, params map[string]interface{}, response interface{}) (status int, err error) {
	url := v.buildURL(uri)
	return v.sendRequest("PUT", url, defaultContentType, params, response)
}

// Delete constructs and synchronously sends an API DELETE request
func (v *VouchedAPIClient) Delete(uri string) (status int, err error) {
	url := v.buildURL(uri)
	return v.sendRequest("DELETE", url, defaultContentType, nil, nil)
}

// Send constructs and synchronously sends a graphql request
func (v *VouchedAPIClient) Send(uri string, request map[string]interface{}, response interface{}) (status int, err error) {
	url := v.buildURL(uri)
	return v.sendRequest("POST", url, defaultContentType, request, response)
}

func (v *VouchedAPIClient) buildURL(uri string) string {
	path := v.Path
	if len(path) == 1 && path == "/" {
		path = ""
	} else if len(path) > 1 && strings.Index(path, "/") != 0 {
		path = fmt.Sprintf("/%s", path)
	}
	return fmt.Sprintf("%s://%s%s/%s", v.Scheme, v.Host, path, uri)
}
