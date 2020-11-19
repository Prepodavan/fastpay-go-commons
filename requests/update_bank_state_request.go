package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

type UpdateBankStateRequest struct {
	TechnicalSignRequest
	Address string           `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State   state_enum.State `json:"state" valid:"required"`
}

func (request *UpdateBankStateRequest) String() string {
	return fmt.Sprintf("updateState%s%v", request.Address, request.State)
}
