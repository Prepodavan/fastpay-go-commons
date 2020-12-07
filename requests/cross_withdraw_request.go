package requests

type CrossWithdrawRequest struct {
	WithdrawRequest
	BankAddress string `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
