package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-category-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/filter-contract-state-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/operation-deal-type-enum"
)

type GetContractPageRequest struct {
	PageSize      int32                                                                     `yaml:"pageSize" json:"pageSize" valid:"required"`
	Bookmark      string                                                                    `yaml:"bookmark" json:"bookmark"`
	EndDate       int64                                                                     `yaml:"endDate" json:"endDate" valid:"range(0|9223372036854775807)"`
	Status        filter_contract_state_enum.FilterContractState                            `yaml:"status" json:"status" valid:"range(0|9223372036854775807)"`
	OperationType operation_deal_type_enum.OperationDealType                                `yaml:"operationType" json:"operationType" valid:"range(0|9223372036854775807)"`
	Address       string                                                                    `yaml:"address" json:"address" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode  int                                                                       `yaml:"currencyCode" json:"currencyCode" valid:"range(0|999)~ErrorCurrencyCodeRange"`
	Category      currency_exchange_contract_category_enum.CurrencyExchangeContractCategory `yaml:"category" json:"category"`
}

func (getContractPage *GetContractPageRequest) SetDefaults() {
	if getContractPage.Category == currency_exchange_contract_category_enum.Undefined {
		getContractPage.Category = currency_exchange_contract_category_enum.CurrencyExchangeContract
	}
}
