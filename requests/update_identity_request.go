package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
)

type UpdateIdentityRequest struct {
	TechnicalSignRequest
	Address      string                          `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	IdentityType identity_type_enum.IdentityType `json:"identityType" valid:"required~ErrorIdentityTypeNotPassed"`
}


func (request *UpdateIdentityRequest) MsgHash() string {
	return fmt.Sprintf("updateIdentity%s%v", request.Address, request.IdentityType)
}
