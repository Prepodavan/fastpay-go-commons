package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-category-enum"

type GetBestRoutesRequest struct {
	Amount             int64                                                                     `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed"`
	CurrencyCodeFrom   int                                                                       `yaml:"currencyCodeFrom" json:"currencyCodeFrom" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CurrencyCodeTo     int                                                                       `yaml:"currencyCodeTo" json:"currencyCodeTo" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CustomerIdentifier string                                                                    `yaml:"customerIdentifier" json:"customerIdentifier" valid:"validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode        string                                                                    `yaml:"countryCode" json:"countryCode" valid:"stringlength(3|3)"`
	To                 string                                                                    `yaml:"to" json:"to" valid:"validHex40or64~ErrorAddressOrIdentifierNotFolowingRegex"`
	Category           currency_exchange_contract_category_enum.CurrencyExchangeContractCategory `yaml:"category" json:"category"`
}

func (getBestRoutes *GetBestRoutesRequest) SetDefaults() {
	if getBestRoutes.Category == currency_exchange_contract_category_enum.Undefined {
		getBestRoutes.Category = currency_exchange_contract_category_enum.CurrencyExchangeContract
	}
}
