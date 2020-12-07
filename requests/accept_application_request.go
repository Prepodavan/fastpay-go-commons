package requests

type AcceptApplicationRequest struct {
	Id              string  `yaml:"id" json:"id" valid:"required,uuid"`
	AddressAcceptor string  `yaml:"addressAcceptor" json:"addressAcceptor" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	MsgHash         string  `yaml:"msgHash" json:"msgHash" valid:"required"`
	Sig             SignDto `yaml:"sig" json:"sig" valid:"required"`
	Exp             int64   `yaml:"exp" json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
