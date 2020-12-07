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
	Address                string                                                      `json:"address"`
	State                  stateEnum.State                                             `json:"state"`
	CurrencyCode           int                                                         `json:"currencyCode"`
	JuridicalType          juridicalEnum.JuridicalType                                 `json:"juridicalType"`
	ClientBankAddress      string                                                      `json:"clientBankAddress"`
	IdentityType           identityEnum.IdentityType                                   `json:"identityType"`
	Owner                  string                                                      `json:"owner"`
	Type                   accountTypeEnum.AccountType                                 `json:"type"`
	AvailableContractTypes []currencyExchangeContractEnum.CurrencyExchangeContractType `json:"availableContractTypes"`
	Identifiers            []string                                                    `json:"identifiers"`
	Encrypted              bool                                                        `json:"encrypted"`
	Created                int64                                                       `json:"created"`
	PublicKey              string                                                      `json:"publicKey"`
	DocType                string                                                      `json:"docType"`
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
