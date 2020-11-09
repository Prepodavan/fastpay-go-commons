package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-category-enum"

type GetBestRoutesRequest struct {
	Amount             int64                                                                       `json:"amount" valid:"required~ErrorAmountNotPassed"`
	CurrencyCodeFrom   int                                                                         `json:"currencyCodeFrom" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CurrencyCodeTo     int                                                                         `json:"currencyCodeTo" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CustomerIdentifier string                                                                      `json:"customerIdentifier" valid:"validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode        string                                                                      `json:"countryCode" valid:"stringlength(3|3)"`
	To                 string                                                                      `json:"to" valid:"validHex40or64~ErrorAddressOrIdentifierNotFolowingRegex"`
	Category           currency_exchange_contract_category_enum.CurrencyExchangeContractCategory `json:"category"`
}

func (getBestRoutes *GetBestRoutesRequest) SetDefaults() {
	if getBestRoutes.Category == currency_exchange_contract_category_enum.Undefined {
		getBestRoutes.Category = currency_exchange_contract_category_enum.CurrencyExchangeContract
	}
}
