package vouched

import "fmt"

// EvaluateFraud evaluates a transaction for payment fraud; see https://edoc.vouched.com
func (v *VouchedAPIClient) EvaluateFraud(params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ReportFraud reports a fraud event; see https://edoc.vouched.com
func (v *VouchedAPIClient) ReportFraud(params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}

// ReportTransaction reports various kinds of transactions including deposits, withdrawals and internal transfer
func (v *VouchedAPIClient) ReportTransaction(txType string, params map[string]interface{}) (interface{}, error) {
	return nil, fmt.Errorf("vouched API method not implemented")
}
