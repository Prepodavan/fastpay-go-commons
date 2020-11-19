package requests

import "fmt"

type AddRegulatorRequest struct {
	TechnicalSignRequest
	Address string `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}

func (request *AddRegulatorRequest) MsgHash() string {
	return fmt.Sprintf("addRegulator%s", request.Address)
}
