package requests

import "fmt"

type SetEnabledByCodeRequest struct {
	TechnicalSignRequest
	Code    int  `json:"code" valid:"required,range(0|9223372036854775807)"`
	Enabled bool `json:"enabled"`
}

func (request *SetEnabledByCodeRequest) MsgHash() string {
	return fmt.Sprintf("setEnabledByCode%v%v", request.Code, request.Enabled)
}
