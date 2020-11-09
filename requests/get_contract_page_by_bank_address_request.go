package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-category-enum"
)

type GetContractPageByBankRequest struct {
	GetContractPageRequest
	BankAddress string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}

func (getContractPage *GetContractPageByBankRequest) SetDefaults() {
	if getContractPage.Category == currency_exchange_contract_category_enum.Undefined {
		getContractPage.Category = currency_exchange_contract_category_enum.CurrencyExchangeContract
	}
}
