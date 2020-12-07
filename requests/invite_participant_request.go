package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/member-deal-type-enum"
)

type InviteParticipantRequest struct {
	Id          string                               `yaml:"id" json:"id" valid:"required,uuid"`
	AddressFrom string                               `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Address     string                               `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40or64~ErrorAddressNotFollowingRegex"`
	MemberType  member_deal_type_enum.MemberDealType `yaml:"memberType" json:"memberType" valid:"required,range(0|2)"`
	MsgHash     string                               `yaml:"msgHash" json:"msgHash" valid:"required"`
	Sig         SignDto                              `yaml:"sig" json:"sig" valid:"required"`
	Exp         int64                                `yaml:"exp" json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
