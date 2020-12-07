package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type CompleteSafeDealDepositRequest struct {
	SafeDealId   string                  `yaml:"safeDealId" json:"safeDealId" valid:"required,uuid"`
	AddressTo    string                  `yaml:"addressAcceptor" json:"addressAcceptor" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyInfo models.CurrencyDealInfo `yaml:"currencyInfo" json:"currencyInfo" valid:"required"`
}
