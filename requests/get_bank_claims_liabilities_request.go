package requests

type GetBankClaimsLiabilitiesRequest struct {
	CurrencyCode int    `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Bank         string `yaml:"bank" json:"bank" valid:"required~ErrorBankAddressNotPassed,validHex40"`
}
