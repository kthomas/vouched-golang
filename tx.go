package vouched

import "fmt"

// VouchedTxTypeDeposit maps to 'transferin' URI; see https://edoc.vouched.com
const VouchedTxTypeDeposit = "transferin"

// VouchedTxTypeWithdrawal maps to 'transferout' URI; see https://edoc.vouched.com
const VouchedTxTypeWithdrawal = "transferout"

// VouchedTxTypeTransfer maps to 'transfer' URI; see https://edoc.vouched.com
const VouchedTxTypeTransfer = "transfer"

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
