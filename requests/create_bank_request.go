package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"

type CreateBankRequest struct {
	Address string           `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State   state_enum.State `yaml:"state" json:"state"`
	Name    string           `yaml:"name" json:"name" valid:"required"`
	BIK     string           `yaml:"bik" json:"bik" valid:"required"`
	MspId   string           `yaml:"mspId" json:"mspId" valid:"required"`
	IsOwner bool             `yaml:"isOwner" json:"isOwner"`
	Conf    string           `yaml:"conf" json:"conf"`
}

func (createBank *CreateBankRequest) SetDefaults() {
	if createBank.State == state_enum.Undefined {
		createBank.State = state_enum.Available
	}
}
