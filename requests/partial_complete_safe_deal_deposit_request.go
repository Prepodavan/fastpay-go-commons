package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type PartialCompleteSafeDealDepositRequest struct {
	SafeDealId   string                  `yaml:"safeDealId" json:"safeDealId" valid:"required,uuid"`
	AddressTo    string                  `yaml:"addressAcceptor" json:"addressAcceptor" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount       int64                   `yaml:"amount" json:"amount" valid:"range(0|9223372036854775807)~ErrorAmountNegative"`
	CurrencyInfo models.CurrencyDealInfo `yaml:"currencyInfo" json:"currencyInfo" valid:"required"`
}
