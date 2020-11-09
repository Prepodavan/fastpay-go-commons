package currency_exchange_contracts_type_enum

type CurrencyExchangeContractsType int

const (
	Undefined CurrencyExchangeContractsType = iota
	CurrencyExchangeConversionContract
	CurrencyExchangeOutgoingCommissionContract
	CurrencyExchangeIncomingCommissionContract
)

func (currencyExchangeContractsType CurrencyExchangeContractsType) Str() string {
	return [...]string{"Undefined", "CurrencyExchangeConversionContract", "CurrencyExchangeOutgoingCommissionContract", "CurrencyExchangeIncomingCommissionContract"}[currencyExchangeContractsType]
}

func Parse(stringCurrencyExchangeContractsType string) CurrencyExchangeContractsType {
	switch stringCurrencyExchangeContractsType {
	case "CurrencyExchangeConversionContract":
		return CurrencyExchangeConversionContract
	case "CurrencyExchangeOutgoingCommissionContract":
		return CurrencyExchangeOutgoingCommissionContract
	case "CurrencyExchangeIncomingCommissionContract":
		return CurrencyExchangeIncomingCommissionContract
	default:
		return Undefined
	}
}

func GetCurrencyExchangeContractsType() []CurrencyExchangeContractsType {
	return []CurrencyExchangeContractsType{Undefined, CurrencyExchangeConversionContract, CurrencyExchangeOutgoingCommissionContract, CurrencyExchangeIncomingCommissionContract}
}
