package currency_exchange_contracts_category_enum

type CurrencyExchangeContractsCategory int

const (
	Undefined CurrencyExchangeContractsCategory = iota
	CurrencyExchangeContract
	OfferExchangeContract
)

func (currencyExchangeContractsCategory CurrencyExchangeContractsCategory) Str() string {
	return [...]string{"Undefined", "CurrencyExchangeContract", "OfferExchangeContract"}[currencyExchangeContractsCategory]
}

func Parse(stringCurrencyExchangeContractsCategory string) CurrencyExchangeContractsCategory {
	switch stringCurrencyExchangeContractsCategory {
	case "CurrencyExchangeContract":
		return CurrencyExchangeContract
	case "OfferExchangeContract":
		return OfferExchangeContract
	default:
		return Undefined
	}
}

func GetCurrencyExchangeContractsCategories() []CurrencyExchangeContractsCategory {
	return []CurrencyExchangeContractsCategory{Undefined, CurrencyExchangeContract, OfferExchangeContract}
}
