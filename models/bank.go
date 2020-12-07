package models

import (
	roleEnum "github.com/SolarLabRU/fastpay-go-commons/enums/access-role-enum"
	typeEnum "github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

type Bank struct {
	Address                string                                  `yaml:"address" yaml:"address" json:"address"`
	Name                   string                                  `yaml:"name" json:"name"`
	BIK                    string                                  `yaml:"bik" json:"bik"`
	State                  state_enum.State                        `yaml:"state" json:"state"`
	CreatedBy              string                                  `yaml:"createdBy" json:"createdBy"`
	Encrypted              bool                                    `yaml:"encrypted" json:"encrypted"`
	MSPId                  string                                  `yaml:"mspId" json:"mspId"`
	Roles                  []roleEnum.AccessRole                   `yaml:"roles" json:"roles"`
	AvailableContractTypes []typeEnum.CurrencyExchangeContractType `yaml:"availableContractTypes" json:"availableContractTypes"`
	Conf                   string                                  `yaml:"conf" json:"conf"`
	DocType                string                                  `yaml:"docType" json:"docType"`
}

func (b *Bank) GetAddress() string {
	return b.Address
}
