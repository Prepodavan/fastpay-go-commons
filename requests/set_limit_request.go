package requests

import "fmt"

type SetLimitRequest struct {
	TechnicalSignRequest
	GetLimitRequest
	Value int64 `yaml:"value" json:"value" valid:"range(0|9223372036854775807)"`
}

func (request *SetLimitRequest) String() string {
	return fmt.Sprintf("setLimit%v%v%v%v%v%v", request.CurrencyCode, request.IdentityType, request.JuridicalType, request.AccountType, request.LimitType, request.Value)
}
