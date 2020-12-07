package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type TopUpDepositRequest struct {
	SafeDealId   string                  `yaml:"safeDealId" json:"safeDealId" valid:"required,uuid"`
	AddressFrom  string                  `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AddressTo    string                  `yaml:"addressTo" json:"addressTo" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyInfo models.CurrencyDealInfo `yaml:"currencyInfo" json:"currencyInfo" valid:"required"`
	Amount       int64                   `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed"`
	NeedAmount   int64                   `yaml:"needAmount" json:"needAmount" valid:"required~ErrorAmountNotPassed"`
}
