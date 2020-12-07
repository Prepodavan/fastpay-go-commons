package requests

import "fmt"

type SetEnabledByCodeRequest struct {
	TechnicalSignRequest
	Code    int  `yaml:"code" json:"code" valid:"required,range(0|9223372036854775807)"`
	Enabled bool `yaml:"enabled" json:"enabled"`
}

func (request *SetEnabledByCodeRequest) String() string {
	return fmt.Sprintf("setEnabledByCode%v%v", request.Code, request.Enabled)
}
