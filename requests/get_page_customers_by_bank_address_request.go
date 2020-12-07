package requests

type GetPageCustomersByBankAddressRequest struct {
	GetPageRequest
	Address string `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
