package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type SaveClaimsRequest struct {
	Claims       []models.ClaimsItem `yaml:"claims" json:"claims" valid:"required"`
	CurrencyCode int                 `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}
