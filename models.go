package vouched

// KYCApplication represents a vouched KYC application
type KYCApplication struct {
	ID        *string                                    `json:"id"`
	Status    *string                                    `json:"status"`
	Submitted *string                                    `json:"submitted"`
	Request   *KYCApplicationRequest                     `json:"request"`
	Result    *KYCApplicationIDVerificationResult        `json:"result"`
	Errors    []*KYCApplicationIDVerificationResultError `json:"errors"`
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
	Errors      []*KYCApplicationIDVerificationResultError     `json:"errors"`
}

// KYCApplicationIDVerificationResultConfidences represents confidence scores for a vouched KYC application id verification response
type KYCApplicationIDVerificationResultConfidences struct {
	ID        *float64 `json:"id"`
	BackID    *float64 `json:"backId"`
	FaceMatch *float64 `json:"faceMatch"`
	IDMatch   *float64 `json:"idMatch"`
	Selfie    *float64 `json:"selfie"`
}

// KYCApplicationIDVerificationResultError represents an error for a vouched KYC application id verification response
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
type KYCApplicationIDVerificationResultError struct {
	Type       *string                                    `json:"type"`
	Message    *string                                    `json:"message"`
	Suggestion *string                                    `json:"suggestion"`
	Errors     []*KYCApplicationIDVerificationResultError `json:"errors"`
}

// KYCApplicationQuery represents a graphql request for a previously-submitted KYC application
type KYCApplicationQuery struct {
	ID         *string `json:"id"`
	Type       *string `json:"type"`
	Status     *string `json:"status"`
	WithPhotos *bool   `json:"withPhotos"`

	// page	Int		Paginate list by page where the page starts at 1, defaults to 1.
	// pageSize	Int		The number of items for a page, max 1000, defaults to 100
	// sortBy	String		Sort the list from ("date", "status").
	// sortOrder	String		Order the sort from ("asc", "desc").
	// to	String		Filter by submitted to ISO8601 date.
	// from	String		Filter by submitted from ISO8601 date.
	// withPhotos	Boolean		Defaults to False. The returned job will contain detailed information idPhoto, idPhotoBack, and userPhoto.
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
