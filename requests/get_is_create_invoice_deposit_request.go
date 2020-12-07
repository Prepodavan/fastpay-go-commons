package requests

type GetIsCreateInvoiceDepositRequest struct {
	Id        string `yaml:"id" json:"id" valid:"required,uuid"`
	AddressTo string `yaml:"addressTo" json:"addressTo" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
