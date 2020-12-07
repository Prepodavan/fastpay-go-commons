package requests

type ConfirmInvitationRequest struct {
	Id               string  `yaml:"id" json:"id" valid:"required,uuid"`
	Address          string  `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount           int64   `yaml:"amount" json:"amount" valid:"range(0|9223372036854775807)"`
	ObligatoryAmount int64   `yaml:"obligatoryAmount" json:"obligatoryAmount" valid:"range(0|9223372036854775807)"`
	MsgHash          string  `yaml:"msgHash" json:"msgHash" valid:"required"`
	Sig              SignDto `yaml:"sig" json:"sig" valid:"required"`
	Exp              int64   `yaml:"exp" json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
