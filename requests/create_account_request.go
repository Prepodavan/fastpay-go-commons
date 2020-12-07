package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/account-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

type CreateAccountRequest struct {
	TechnicalSignRequest
	Address       string                            `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State         state_enum.State                  `yaml:"state" json:"state"`
	CurrencyCode  int                               `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)"`
	JuridicalType juridical_type_enum.JuridicalType `yaml:"juridicalType" json:"juridicalType"`
	IdentityType  identity_type_enum.IdentityType   `yaml:"identityType" json:"identityType"`
	Type          account_type_enum.AccountType     `yaml:"type" json:"type"`
	Identifiers   []string                          `yaml:"identifiers" json:"identifiers" valid:"validHex64~ErrorIdentifierNotFolowingRegex"`
	PublicKey     string                            `yaml:"publicKey" json:"publicKey"`
	Owner         string                            `yaml:"owner" json:"owner"`
	MsgHash       string                            `yaml:"msgHash" json:"msgHash"`
	Sig           SignDto                           `yaml:"sig" json:"sig"`
}

func (createAccount *CreateAccountRequest) SetDefaults() {
	if createAccount.IdentityType == identity_type_enum.Undefined {
		createAccount.IdentityType = identity_type_enum.None
	}
	if createAccount.State == state_enum.Undefined {
		createAccount.State = state_enum.Available
	}
	if createAccount.Type == account_type_enum.Undefined {
		createAccount.Type = account_type_enum.Client
	}
	if createAccount.Identifiers == nil {
		createAccount.Identifiers = []string{}
	}
}
