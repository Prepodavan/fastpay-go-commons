package requests

import "fmt"

type AddRegulatorRequest struct {
	TechnicalSignRequest
	Address string `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}

func (request *AddRegulatorRequest) GetAddress() string {
	return request.Address
}

func (request *AddRegulatorRequest) String() string {
	return fmt.Sprintf("addRegulator%s", request.Address)
}
