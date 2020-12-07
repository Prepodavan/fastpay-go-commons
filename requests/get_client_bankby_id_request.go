package requests

type GetClientBankByAddressRequest struct {
	Address string `yaml:"address" json:"address" valid:"required~ErrorBankAddressNotPassed"`
}
