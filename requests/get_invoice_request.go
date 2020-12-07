package requests

type GetInvoiceRequest struct {
	Number    string `yaml:"number" json:"number" valid:"required~ErrorNumberNotPassed,maxstringlength(255)"`
	Recipient string `yaml:"recipient" json:"recipient" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
