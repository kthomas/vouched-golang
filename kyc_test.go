package vouched

import (
	"testing"
)

func TestKYCSubmitApplication(t *testing.T) {
	apiClient, _ := NewVouchedAPIClient()
	resp, err := apiClient.SubmitApplication(map[string]interface{}{
		"type": "id-verification",
		// "callbackURL": "https://google.com/gdpr",
		"parameters": map[string]interface{}{},
	})
	if err != nil {
		t.Errorf("KYC application submission failed; %s", err.Error())
		return
	}
	log.Debugf("KYC application submission response: %v", resp.(*KYCApplicationIDVerificationResult).Errors[0])
}

func TestKYCGetApplication(t *testing.T) {
	apiClient, _ := NewVouchedAPIClient()
	resp, err := apiClient.GetApplication("does-not-exist")
	if err != nil {
		t.Errorf("KYC application (job) query failed; %s", err.Error())
		return
	}
	log.Debugf("KYC application query response: %v", resp)
}
