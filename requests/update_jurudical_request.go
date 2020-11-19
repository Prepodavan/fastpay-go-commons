package requests

import (
	"fmt"
	juridical_type_enum "github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
)

type UpdateJuridicalRequest struct {
	TechnicalSignRequest
	Address       string                            `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	JuridicalType juridical_type_enum.JuridicalType `json:"juridicalType" valid:"required~ErrorJuridicalTypeNotPassed"`
}

func (request *UpdateJuridicalRequest) String() string {
	return fmt.Sprintf("updateJuridical%s%v", request.Address, request.JuridicalType)
}
