package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"

type UpdateAvailableContractTypes struct {
	TechnicalSignRequest
	Address                string                                                              `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AvailableContractTypes []currency_exchange_contract_type_enum.CurrencyExchangeContractType `json:"availableContractTypes"`
}
