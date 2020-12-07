package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

type UpdateBankStateRequest struct {
	TechnicalSignRequest
	Address string           `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State   state_enum.State `yaml:"state" json:"state" valid:"required"`
}

func (request *UpdateBankStateRequest) GetAddress() string {
	return request.Address
}

func (request *UpdateBankStateRequest) String() string {
	return fmt.Sprintf("updateState%s%v", request.Address, request.State)
}
