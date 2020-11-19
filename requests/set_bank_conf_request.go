package requests

import "fmt"

type SetBankConfRequest struct {
	TechnicalSignRequest
	Conf string `json:"conf" valid:"required"`
}

func (request *SetBankConfRequest) String() string {
	return fmt.Sprintf("setBankConf%s", request.Conf)
}