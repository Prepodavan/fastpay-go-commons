package requests

type CrossWithdrawConfirmRequest struct {
	WithdrawConfirmRequest
	BankAddress string `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
