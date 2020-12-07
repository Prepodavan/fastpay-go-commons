package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ContractUpdateRequest struct {
	models.CurrencyExchangeContractMutable
	BankAddress string  `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
	Sig         SignDto `yaml:"sig" json:"sig"`
	MsgHash     string  `yaml:"msgHash" json:"msgHash"`
	Exp         int64   `yaml:"exp" json:"exp"`
}
