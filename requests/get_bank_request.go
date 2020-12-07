package requests

type GetBankRequest struct {
	MSPId   string `yaml:"mspId" json:"mspId" valid:"required"`
	Address string `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
