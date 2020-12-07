package requests

type CrossWithdrawRejectRequest struct {
	WithdrawRejectRequest
	BankAddress string `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
