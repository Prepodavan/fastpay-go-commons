package currency_exchange_contract_type_enum

type CurrencyExchangeContractType int

const (
	Undefined CurrencyExchangeContractType = iota
	CurrencyExchangeConversionContract
	CurrencyExchangeOutgoingCommissionContract
	CurrencyExchangeIncomingCommissionContract
)

func (currencyExchangeContractsType CurrencyExchangeContractType) Str() string {
	return [...]string{"Undefined", "CurrencyExchangeConversionContract", "CurrencyExchangeOutgoingCommissionContract", "CurrencyExchangeIncomingCommissionContract"}[currencyExchangeContractsType]
}

func Parse(stringCurrencyExchangeContractsType string) CurrencyExchangeContractType {
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

func GetCurrencyExchangeContractsType() []CurrencyExchangeContractType {
	return []CurrencyExchangeContractType{Undefined, CurrencyExchangeConversionContract, CurrencyExchangeOutgoingCommissionContract, CurrencyExchangeIncomingCommissionContract}
}
