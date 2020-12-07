package requests

import "fmt"

type SetIssueLimitRequest struct {
	TechnicalSignRequest
	Address      string `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Value        int64  `yaml:"value" json:"value" valid:"range(0|9223372036854775807)"`
	CurrencyCode int    `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}

func (request *SetIssueLimitRequest) String() string {
	return fmt.Sprintf("setIssueLimitBank%v%s%v", request.CurrencyCode, request.Address, request.Value)
}
