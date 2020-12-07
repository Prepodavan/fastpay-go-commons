package requests

type GetPageByPayerRequest struct {
	GetPageRequest
	Payer string `yaml:"payer" json:"payer" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
