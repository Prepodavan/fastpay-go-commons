package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ContractUpdateRequest struct {
	models.CurrencyExchangeContractMutable
	BankAddress string  `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
	Sig         SignDto `json:"sig"`
	MsgHash     string  `json:"msgHash"`
}
