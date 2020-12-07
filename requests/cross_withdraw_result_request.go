package requests

type CrossWithdrawResultRequest struct {
	WithdrawResultRequest
	CurrencyCode int    `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	BankAddress  string `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
