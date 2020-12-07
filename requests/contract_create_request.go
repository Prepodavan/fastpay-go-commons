package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-category-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type ContractCreateRequest struct {
	Id                   string                                                                    `yaml:"id" json:"id" valid:"required"`
	AddressAccountSell   string                                                                    `yaml:"addressAccountSell" json:"addressAccountSell" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	AddressAccountBuy    string                                                                    `yaml:"addressAccountBuy" json:"addressAccountBuy" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	AddressCommission    string                                                                    `yaml:"addressCommission" json:"addressCommission" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyInfoSell     models.CurrencyInfo                                                       `yaml:"currencyInfoSell" json:"currencyInfoSell"`
	CurrencyInfoBuy      models.CurrencyInfo                                                       `yaml:"currencyInfoBuy" json:"currencyInfoBuy"`
	Category             currency_exchange_contract_category_enum.CurrencyExchangeContractCategory `yaml:"category" json:"category" valid:"required"`
	Type                 currency_exchange_contract_type_enum.CurrencyExchangeContractType         `yaml:"type" json:"type" valid:"range(0|3)"`
	Price                float64                                                                   `yaml:"price" json:"price" valid:"required~ErrorAmountNotPassed,range(0|9223372036854775807)"`
	FractionalCommission float64                                                                   `yaml:"fractionalCommission" json:"fractionalCommission" valid:"range(0|1)"`
	MaxCommission        int64                                                                     `yaml:"maxCommission" json:"maxCommission" valid:"range(0|9223372036854775807)"`
	MinAmount            int64                                                                     `yaml:"minAmount" json:"minAmount" valid:"range(0|9223372036854775807)~ErrorAmountNegative"`
	MaxAmount            int64                                                                     `yaml:"maxAmount" json:"maxAmount" valid:"required~ErrorAmountNotPassed,range(0|9223372036854775807)~ErrorAmountNegative"`
	StartDate            int64                                                                     `yaml:"startDate" json:"startDate" valid:"required~ErrorTimestampNotPassed,range(0|9223372036854775807)~ErrorTimestampNegative"`
	EndDate              int64                                                                     `yaml:"endDate" json:"endDate" valid:"required~ErrorTimestampNotPassed,range(0|9223372036854775807)~ErrorTimestampNegative"`
	BankAddress          string                                                                    `yaml:"bankAddress" json:"bankAddress"`
	Sig                  SignDto                                                                   `yaml:"sig" json:"sig"`
	MsgHash              string                                                                    `yaml:"msgHash" json:"msgHash"`
	Exp                  int64                                                                     `yaml:"exp" json:"exp"`
}

func (createContract *ContractCreateRequest) SetDefaults() {
	if createContract.Category == currency_exchange_contract_category_enum.Undefined {
		createContract.Category = currency_exchange_contract_category_enum.CurrencyExchangeContract
	}
}
