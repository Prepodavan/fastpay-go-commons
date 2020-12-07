package requests

import "fmt"

type ExecuteClearingRequest struct {
	TechnicalSignRequest
	CurrencyCode int `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}

func (request *ExecuteClearingRequest) String() string {
	return fmt.Sprintf("execute%v", request.CurrencyCode)
}
