package requests

import "fmt"

type SetBankConfRequest struct {
	TechnicalSignRequest
	Conf string `json:"conf" valid:"required"`
}

func (request *SetBankConfRequest) MsgHash() string {
	return fmt.Sprintf("setBankConf%s", request.Conf)
}