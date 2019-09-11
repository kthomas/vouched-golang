package vouched

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func readImage(path string) *string {
	img, err := ioutil.ReadFile(path)
	if err != nil {
		log.Warningf("Failed to read image from path %s", path)
		return nil
	}

	imgBase64 := base64.StdEncoding.EncodeToString(img)
	log.Debugf("Read %d-byte image from path %s (%d-bytes, base64 encoded)", len(img), path, len(imgBase64))
	return stringOrNil(imgBase64)
}

func submitValidKYCApplication() (*KYCApplicationResponse, error) {
	userPhotoBase64 := readImage("test/data/selfie.jpg")
	idPhotoBase64 := readImage("test/data/id.jpg")

	apiClient, _ := NewVouchedAPIClient()
	resp, err := apiClient.SubmitApplication(map[string]interface{}{
		"type":        "id-verification",
		"callbackURL": "https://google.com/gdpr",
		"params": map[string]interface{}{
			"firstName": "Janice",
			"lastName":  "Way",
			"dob":       "06/22/1990",
			"userPhoto": userPhotoBase64,
			"idPhoto":   idPhotoBase64,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.(*KYCApplicationResponse), nil
}

func TestKYCSubmitApplicationIncompleteRequest(t *testing.T) {
	apiClient, _ := NewVouchedAPIClient()
	resp, err := apiClient.SubmitApplication(map[string]interface{}{
		"type": "id-verification",
		// "callbackURL": "https://google.com/gdpr",
		"params": map[string]interface{}{},
	})
	if err != nil {
		t.Errorf("KYC application submission failed; %s", err.Error())
		return
	}
	if len(resp.(*KYCApplicationResponse).Errors) == 0 {
		t.Fail()
	}
}

func TestKYCSubmitApplicationValidID(t *testing.T) {
	userPhotoBase64 := readImage("test/data/selfie.jpg")
	idPhotoBase64 := readImage("test/data/id.jpg")

	apiClient, _ := NewVouchedAPIClient()
	resp, err := apiClient.SubmitApplication(map[string]interface{}{
		"type":        "id-verification",
		"callbackURL": "https://google.com/gdpr",
		"params": map[string]interface{}{
			"firstName": "Janice",
			"lastName":  "Way",
			"dob":       "06/22/1990",
			"userPhoto": userPhotoBase64,
			"idPhoto":   idPhotoBase64,
		},
	})
	if err != nil {
		t.Errorf("KYC application submission failed; %s", err.Error())
		return
	}
	log.Debugf("KYC application submission response: %v", *resp.(*KYCApplicationResponse).Data.Job.ID)
}

func TestKYCGetApplicationInvalidIdentifier(t *testing.T) {
	apiClient, _ := NewVouchedAPIClient()
	resp, err := apiClient.GetApplication("does-not-exist")
	if err != nil {
		t.Errorf("KYC application (job) query failed; %s", err.Error())
		return
	}
	if resp != nil {
		t.Fail()
	}
}

func TestKYCGetApplicationValidIdentifier(t *testing.T) {
	kycSubmissionResp, err := submitValidKYCApplication()
	if err != nil {
		t.Errorf("KYC application (job) submission failed; %s", err.Error())
		return
	}

	apiClient, _ := NewVouchedAPIClient()
	resp, err := apiClient.GetApplication(*kycSubmissionResp.Data.Job.ID)
	if err != nil {
		t.Errorf("KYC application (job) query failed; %s", err.Error())
		return
	}
	if resp == nil {
		t.Fail()
	}
	if *kycSubmissionResp.Data.Job.ID != *resp.(*KYCApplication).ID {
		t.Fail()
	}
	log.Debugf("KYC application query response job ID: %s", *resp.(*KYCApplication).ID)

}
