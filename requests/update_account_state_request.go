package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

type UpdateAccountStateRequest struct {
	TechnicalSignRequest
	Address string           `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State   state_enum.State `json:"state" valid:"required"`
}

func (updateAccountStateRequest *UpdateAccountStateRequest) MsgHash() string {
	return fmt.Sprintf("%s%v", updateAccountStateRequest.Address, updateAccountStateRequest.State)
}
