package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"
	"strconv"
)

type UpdateAvailableContractTypes struct {
	TechnicalSignRequest
	Address                string                                                              `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AvailableContractTypes []currency_exchange_contract_type_enum.CurrencyExchangeContractType `json:"availableContractTypes"`
}

func (request *UpdateAvailableContractTypes) GetAddress() string {
	return request.Address
}

func (request *UpdateAvailableContractTypes) String() string {

	concatenatedTypes := ""

	for _, v := range request.AvailableContractTypes {
		concatenatedTypes += strconv.Itoa(int(v))
	}

	return fmt.Sprintf("updateAvailableContractTypes%s%s", request.Address, concatenatedTypes)
}
