package requests

type CrossWithdrawConfirmRequest struct {
	WithdrawConfirmRequest
	BankId string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}
