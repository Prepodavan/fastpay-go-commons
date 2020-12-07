package models

import (
	accountTypeEnum "github.com/SolarLabRU/fastpay-go-commons/enums/account-type-enum"
	currencyExchangeContractEnum "github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"
	identityEnum "github.com/SolarLabRU/fastpay-go-commons/enums/identity-type-enum"
	juridicalEnum "github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type-enum"
	stateEnum "github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
)

// +k8s:deepcopy-gen=true
// +gen * set ring linkedlist slice:"Any,Where,Count,Select[AddressOfAccount],Aggregate[AddressOfAccount],All"
type Account struct {
	Address                string                                                      `yaml:"address" json:"address"`
	State                  stateEnum.State                                             `yaml:"state" json:"state"`
	CurrencyCode           int                                                         `yaml:"currencyCode" json:"currencyCode"`
	JuridicalType          juridicalEnum.JuridicalType                                 `yaml:"juridicalType" json:"juridicalType"`
	ClientBankAddress      string                                                      `yaml:"clientBankAddress" json:"clientBankAddress"`
	IdentityType           identityEnum.IdentityType                                   `yaml:"identityType" json:"identityType"`
	Owner                  string                                                      `yaml:"owner" json:"owner"`
	Type                   accountTypeEnum.AccountType                                 `yaml:"type" json:"type"`
	AvailableContractTypes []currencyExchangeContractEnum.CurrencyExchangeContractType `yaml:"availableContractTypes" json:"availableContractTypes"`
	Identifiers            []string                                                    `yaml:"identifiers" json:"identifiers"`
	Encrypted              bool                                                        `yaml:"encrypted" json:"encrypted"`
	Created                int64                                                       `yaml:"created" json:"created"`
	PublicKey              string                                                      `yaml:"publicKey" json:"publicKey"`
	DocType                string                                                      `yaml:"docType" json:"docType"`
}

func (in *Account) Hash() string {
	return in.Address
}

type AccountsLoader interface {
	Account(string) (*Account, error)
}

func AccountLoaderAggregator(loader AccountsLoader, err *error) func(AccountSlice, AddressOfAccount) AccountSlice {
	return func(acc AccountSlice, addr AddressOfAccount) AccountSlice {
		if *err != nil {
			return nil
		}
		var account *Account
		account, *err = loader.Account(string(addr))
		return append(acc, account)
	}
}

// +gen set ring linkedlist slice:"Distinct,Any,All,Where,Count,Aggregate[AccountSlice],Select[*Account]"
type AddressOfAccount string
