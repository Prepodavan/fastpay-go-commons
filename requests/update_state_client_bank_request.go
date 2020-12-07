package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

type UpdateStateClientBankRequest struct {
	TechnicalSignRequest
	State   state_enum.State `yaml:"state" json:"state" valid:"required"`
	Address string           `yaml:"address" json:"address" valid:"required~ErrorBankAddressNotPassed"`
}

func (request *UpdateStateClientBankRequest) String() string {
	return fmt.Sprintf("updateState%s%v", request.Address, request.State)
}
