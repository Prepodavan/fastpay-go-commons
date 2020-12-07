package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/account-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/limit-type-enum"
)

type GetLimitRequest struct {
	LimitType     limit_type_enum.LimitType         `yaml:"limitType" json:"limitType" valid:"required,range(1|4)~ErrorLimitTypeIncorrect"`
	IdentityType  identity_type_enum.IdentityType   `yaml:"identityType" json:"identityType" valid:"required,range(1|3)~ErrorIdentityTypeIncorrect"`
	JuridicalType juridical_type_enum.JuridicalType `yaml:"juridicalType" json:"juridicalType" valid:"range(0|3)~ErrorJuridicalTypeIncorrect"`
	AccountType   account_type_enum.AccountType     `yaml:"accountType" json:"accountType" valid:"required,range(1|4)~ErrorAccountTypeIncorrect"`
	CurrencyCode  int                               `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}
