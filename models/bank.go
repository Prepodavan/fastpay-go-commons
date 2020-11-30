package models

import (
	roleEnum "github.com/SolarLabRU/fastpay-go-commons/enums/access-role-enum"
	typeEnum "github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

type Bank struct {
	Address                string                                  `json:"address"`
	Name                   string                                  `json:"name"`
	BIK                    string                                  `json:"bik"`
	State                  state_enum.State                        `json:"state"`
	CreatedBy              string                                  `json:"createdBy"`
	Encrypted              bool                                    `json:"encrypted"`
	MSPId                  string                                  `json:"mspId"`
	Roles                  []roleEnum.AccessRole                   `json:"roles"`
	AvailableContractTypes []typeEnum.CurrencyExchangeContractType `json:"availableContractTypes"`
	Conf                   string                                  `json:"conf"`
	DocType                string                                  `json:"docType"`
}

func (b *Bank) GetAddress() string {
	return b.Address
}
