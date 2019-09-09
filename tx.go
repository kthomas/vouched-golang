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
	var resp map[string]interface{}
	status, err := v.Post(fmt.Sprintf("im/transaction?graphScoreResponse=false"), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to evaluate tx for payment fraud via vouched API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ReportFraud reports a fraud event; see https://edoc.vouched.com
func (v *VouchedAPIClient) ReportFraud(params map[string]interface{}) (interface{}, error) {
	var resp map[string]interface{}
	status, err := v.Post("im/admin/jax/feg", params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to report fraud event via vouched API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}

// ReportTransaction reports various kinds of transactions including deposits, withdrawals and internal transfer
func (v *VouchedAPIClient) ReportTransaction(txType string, params map[string]interface{}) (interface{}, error) {
	if txType != VouchedTxTypeDeposit && txType != VouchedTxTypeWithdrawal && txType != VouchedTxTypeTransfer {
		return nil, fmt.Errorf("Invalid tx type provided: %s", txType)
	}
	var resp map[string]interface{}
	status, err := v.Post(fmt.Sprintf("im/account/%s?graphScoreResponse=false", txType), params, &resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to report tx via vouched API; status: %d; %s", status, err.Error())
	}
	return resp, nil
}
