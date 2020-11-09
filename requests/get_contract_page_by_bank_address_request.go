package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contracts-category-enum"
)

type GetContractPageByBankRequest struct {
	GetContractPageRequest
	BankAddress string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}

func (getContractPage *GetContractPageByBankRequest) SetDefaults() {
	if getContractPage.Type == currency_exchange_contracts_category_enum.Undefined {
		getContractPage.Type = currency_exchange_contracts_category_enum.CurrencyExchangeContract
	}
}
