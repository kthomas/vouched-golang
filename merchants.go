package vouched

import "fmt"

// Merchant aggregation

// CreateMerchant creates a merchant account
func (v *VouchedAPIClient) CreateMerchant(params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// GetMerchant creates a merchant account
func (v *VouchedAPIClient) GetMerchant(merchantID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UpdateMerchant creates a merchant account
func (v *VouchedAPIClient) UpdateMerchant(merchantID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// Merchant KYC

// GetMerchantApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) GetMerchantApplication(applicationID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// SubmitMerchantApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) SubmitMerchantApplication(params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ListMerchantApplicationDocuments see https://edoc.vouched.com
func (v *VouchedAPIClient) ListMerchantApplicationDocuments(applicationID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// DownloadMerchantApplicationDocument see https://edoc.vouched.com
func (v *VouchedAPIClient) DownloadMerchantApplicationDocument(applicationID, documentID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UploadMerchantApplicationDocument see https://edoc.vouched.com
func (v *VouchedAPIClient) UploadMerchantApplicationDocument(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UploadMerchantApplicationDocumentVerificationImage see https://edoc.vouched.com
func (v *VouchedAPIClient) UploadMerchantApplicationDocumentVerificationImage(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ApproveMerchantApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) ApproveMerchantApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// RejectMerchantApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) RejectMerchantApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UndecideMerchantApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) UndecideMerchantApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ProvideMerchantApplicationResponse see https://edoc.vouched.com
func (v *VouchedAPIClient) ProvideMerchantApplicationResponse(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// Merchant KYB

// RejectMerchantBusinessApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) RejectMerchantBusinessApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UndecideMerchantBusinessApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) UndecideMerchantBusinessApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// GetMerchantBusinessApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) GetMerchantBusinessApplication(applicationID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ReevaluateMerchantBusinessApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) ReevaluateMerchantBusinessApplication(applicationID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// SubmitMerchantBusinessApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) SubmitMerchantBusinessApplication(merchantID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ListMerchantBusinessApplicationDocuments see https://edoc.vouched.com
func (v *VouchedAPIClient) ListMerchantBusinessApplicationDocuments(applicationID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// DownloadMerchantBusinessApplicationDocument see https://edoc.vouched.com
func (v *VouchedAPIClient) DownloadMerchantBusinessApplicationDocument(applicationID, documentID string) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UploadMerchantBusinessApplicationDocument see https://edoc.vouched.com
func (v *VouchedAPIClient) UploadMerchantBusinessApplicationDocument(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// UploadMerchantBusinessApplicationDocumentVerificationImage see https://edoc.vouched.com
func (v *VouchedAPIClient) UploadMerchantBusinessApplicationDocumentVerificationImage(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ApproveMerchantBusinessApplication see https://edoc.vouched.com
func (v *VouchedAPIClient) ApproveMerchantBusinessApplication(applicationID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// EvaluateMerchantFraud evaluates a transaction for payment fraud on behalf of a given merchant; see https://edoc.vouched.com
func (v *VouchedAPIClient) EvaluateMerchantFraud(merchantID string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ReportMerchantTransaction reports various kinds of transactions including deposits, withdrawals and internal transfer
func (v *VouchedAPIClient) ReportMerchantTransaction(merchantID, txType string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}
