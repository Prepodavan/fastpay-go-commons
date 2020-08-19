package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type SaveClaimsRequest struct {
	Claims       []models.ClaimsItem `json:"claims" valid:"required"`
	CurrencyCode int                 `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
}
