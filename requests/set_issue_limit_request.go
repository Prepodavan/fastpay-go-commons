package requests

type SetIssueLimitRequest struct {
	Address      string `json:"address" validate:"required,validHex40"`
	Value        int64  `json:"value" validate:"gte=0"`
	CurrencyCode int    `json:"currencyCode" validate:"required,gte=0,lte=999"`
}