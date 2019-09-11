package vouched

// KYCApplication represents a vouched KYC application
type KYCApplication struct {
	ID        *string                             `json:"id"`
	Status    *string                             `json:"status"`
	Submitted *string                             `json:"submitted"`
	Request   *KYCApplicationRequest              `json:"request"`
	Result    *KYCApplicationIDVerificationResult `json:"result"`
	Errors    []*Error                            `json:"errors"`
}

// KYCApplicationRequestParameters represents a vouched KYC job request
type KYCApplicationRequestParameters struct {
	FirstName         *string `json:"firstName"`
	LastName          *string `json:"lastName"`
	DOB               *string `json:"dob"`
	UserPhoto         *string `json:"userPhoto"`
	IDPhoto           *string `json:"idPhoto"`
	IDPhotoBack       *string `json:"idPhotoBack"`
	TwicPhoto         *string `json:"twicPhoto"`
	CarInsurancePhoto *string `json:"carInsurancePhoto"`
	DOTPhoto          *string `json:"dotPhoto"`
}

// KYCApplicationRequest represents a vouched KYC application request
type KYCApplicationRequest struct {
	Type        *string                          `json:"type"`
	CallbackURL *string                          `json:"callbackURL"`
	Parameters  *KYCApplicationRequestParameters `json:"parameters"`
}

// KYCApplicationResponse represents a vouched KYC application submission response
type KYCApplicationResponse struct {
	Data   *KYCApplicationResponseData `json:"data"`
	Errors []*Error                    `json:"errors"`
}

// KYCApplicationResponseData represents a vouched KYC application response "data" object
type KYCApplicationResponseData struct {
	Job *KYCApplication `json:"submitJob"`
}

// KYCApplicationIDVerificationResult represents a vouched KYC application id verification response
type KYCApplicationIDVerificationResult struct {
	ID          *string                                        `json:"id"`
	Type        *string                                        `json:"type"`
	FirstName   *string                                        `json:"firstName"`
	LastName    *string                                        `json:"lastName"`
	State       *string                                        `json:"state"`
	Country     *string                                        `json:"country"`
	DOB         *string                                        `json:"dob"`
	Success     *bool                                          `json:"success"`
	Confidences *KYCApplicationIDVerificationResultConfidences `json:"confidences"`
	Errors      []*Error                                       `json:"errors"`
}

// KYCApplicationIDVerificationResultConfidences represents confidence scores for a vouched KYC application id verification response
type KYCApplicationIDVerificationResultConfidences struct {
	ID        *float64 `json:"id"`
	BackID    *float64 `json:"backId"`
	FaceMatch *float64 `json:"faceMatch"`
	IDMatch   *float64 `json:"idMatch"`
	Selfie    *float64 `json:"selfie"`
}

// Error represents an error for a vouched KYC API call or id verification response
// The following are valid error types:
// 		- InvalidRequestError
// 		- FaceMatchError
// 		- NameMatchError
// 		- BirthDateMatchError
// 		- InvalidIdPhotoError
// 		- InvalidUserPhotoError
// 		- UNAUTHENTICATED/AuthenticationError
// 		- ConnectionError
// 		- InvalidRequestError
// 		- InvalidIdBackPhotoError
// 		- UnknownSystemError
// 		- InvalidIdError
type Error struct {
	Type       *string `json:"type"`
	Message    *string `json:"message"`
	Suggestion *string `json:"suggestion"`
	// Errors     []*Error `json:"errors"`
}

// KYCApplicationQuery represents a graphql request for a previously-submitted KYC application
type KYCApplicationQuery struct {
	ID         *string `json:"id"`
	Type       *string `json:"type"`
	Status     *string `json:"status"`
	WithPhotos *bool   `json:"withPhotos"`
}

// KYCApplicationQueryResponse represents a vouched KYC application query response
type KYCApplicationQueryResponse struct {
	Data   *KYCApplicationQueryResponseData `json:"data"`
	Errors []*Error                         `json:"errors"`
}

// KYCApplicationQueryResponseData represents a vouched KYC application query response "data" object
type KYCApplicationQueryResponseData struct {
	Jobs *KYCApplicationResponseJobs `json:"jobs"`
}

// KYCApplicationResponseJobs represents a vouched KYC application response data "jobs" object
type KYCApplicationResponseJobs struct {
	Items []*KYCApplication `json:"items"`
}

// IsAccepted returns true if the KYC application has been accepted
func (k *KYCApplication) IsAccepted() bool {
	return k.Status != nil && *k.Status == "completed" && k.Result != nil && k.Result.Success != nil && *k.Result.Success && len(k.Result.Errors) == 0
}

// IsRejected returns true if the KYC application has been rejected
func (k *KYCApplication) IsRejected() bool {
	return k.Status != nil && *k.Status == "completed" && k.Result != nil && k.Result.Success != nil && (!*k.Result.Success || len(k.Result.Errors) > 0)
}

// IsUnderReview returns true if the KYC application is currently pending review
func (k *KYCApplication) IsUnderReview() bool {
	return k.Status != nil && *k.Status == "active"
}
