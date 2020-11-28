package requests

type GetIsCreateInvoiceDepositRequest struct {
	Id        string `json:"id" valid:"required,uuid"`
	AddressTo string `json:"addressTo" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
