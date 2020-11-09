package currency_exchange_contract_category_enum

type CurrencyExchangeContractCategory int

const (
	Undefined CurrencyExchangeContractCategory = iota
	CurrencyExchangeContract
	OfferExchangeContract
)

func (currencyExchangeContractsCategory CurrencyExchangeContractCategory) Str() string {
	return [...]string{"Undefined", "CurrencyExchangeContract", "OfferExchangeContract"}[currencyExchangeContractsCategory]
}

func Parse(stringCurrencyExchangeContractsCategory string) CurrencyExchangeContractCategory {
	switch stringCurrencyExchangeContractsCategory {
	case "CurrencyExchangeContract":
		return CurrencyExchangeContract
	case "OfferExchangeContract":
		return OfferExchangeContract
	default:
		return Undefined
	}
}

func GetCurrencyExchangeContractsCategories() []CurrencyExchangeContractCategory {
	return []CurrencyExchangeContractCategory{Undefined, CurrencyExchangeContract, OfferExchangeContract}
}
