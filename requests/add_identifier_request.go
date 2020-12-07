package requests

type AddIdentifierRequest struct {
	Address    string  `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Identifier string  `yaml:"identifier" json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64"`
	MsgHash    string  `yaml:"msgHash" json:"msgHash" valid:"required"`
	Sig        SignDto `yaml:"sig" json:"sig" valid:"required"`
}
