package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"
)

type UpdateAvailableContractTypes struct {
	TechnicalSignRequest
	Address                string                                                              `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AvailableContractTypes []currency_exchange_contract_type_enum.CurrencyExchangeContractType `json:"availableContractTypes"`
}

func (request *UpdateAvailableContractTypes) String() string {

	concatenatedTypes := ""

	for _, v := range request.AvailableContractTypes {
		concatenatedTypes += string(int(v))
	}

	return fmt.Sprintf("updateAvailableContractTypes%s%s", request.Address, concatenatedTypes)
}
