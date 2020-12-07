package requests

type GetAllClaimsRequest struct {
	CurrencyCode int `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}
