package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"
)

type CreateCurrencyRequest struct {
	Code       int                             `yaml:"code" json:"code" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Name       string                          `yaml:"name" json:"name" valid:"required"`
	Type       currency_type_enum.CurrencyType `yaml:"type" json:"type"`
	Unit       string                          `yaml:"unit" json:"unit"`
	Symbol     string                          `yaml:"symbol" json:"symbol" valid:"required~ErrorSymbolNotPassed,stringlength(3|3)"`
	Decimals   int                             `yaml:"decimals" json:"decimals" valid:"required~ErrorDecimalsNotPassed,range(0|8)~ErrorDecimalsRange"`
	Properties map[string]string               `yaml:"properties" json:"properties"`
	Enabled    bool                            `yaml:"enabled" json:"enabled"`
}

func (createCurrencyRequest *CreateCurrencyRequest) SetDefaults() {
	if createCurrencyRequest.Type == currency_type_enum.Undefined {
		createCurrencyRequest.Type = currency_type_enum.DigitalCurrency
	}
	if createCurrencyRequest.Properties == nil {
		createCurrencyRequest.Properties = make(map[string]string)
	}
}
