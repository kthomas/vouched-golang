package vouched

import "fmt"

// KYC

// GetApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) GetApplication(applicationID string) (interface{}, error) {
	resp := &KYCApplicationQueryResponse{}
	status, err := v.Post(graphqlQueryJobs, map[string]interface{}{"id": applicationID, "withPhotos": false}, resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve KYC application via vouched API; status: %d; %s", status, err.Error())
	}
	if resp != nil && resp.Data != nil && resp.Data.Jobs != nil && resp.Data.Jobs.Items != nil && len(resp.Data.Jobs.Items) == 1 {
		return resp.Data.Jobs.Items[0], nil
	}
	return nil, nil
}

// SubmitApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) SubmitApplication(params map[string]interface{}) (interface{}, error) {
	if _, callbackURLOk := params["callbackURL"].(string); !callbackURLOk {
		params["callbackURL"] = vouchedCallbackURL
	}
	resp := &KYCApplicationResponse{}
	status, err := v.Post(graphqlSubmitJobMutation, params, resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload consumer KYC document via vouched API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ProvideApplicationResponse see https://edoc.vouched.com
func (v *VouchedAPIClient) ProvideApplicationResponse(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ListApplicationDocuments see https://edoc.vouched.com
func (v *VouchedAPIClient) ListApplicationDocuments(applicationID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// DownloadApplicationDocument see https://edoc.vouched.com
func (v *VouchedAPIClient) DownloadApplicationDocument(applicationID, documentID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UploadApplicationDocument see https://edoc.vouched.com
func (v *VouchedAPIClient) UploadApplicationDocument(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UploadApplicationDocumentVerificationImage see https://edoc.vouched.com
func (v *VouchedAPIClient) UploadApplicationDocumentVerificationImage(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ApproveApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) ApproveApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// RejectApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) RejectApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UndecideApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) UndecideApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}
