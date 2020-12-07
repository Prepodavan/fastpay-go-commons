package requests

type GetListByPayerRequest struct {
	Payer string `yaml:"payer" json:"payer" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
