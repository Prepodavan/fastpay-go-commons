package models

import "github.com/SolarLabRU/fastpay-go-commons/enums/state"
import "github.com/SolarLabRU/fastpay-go-commons/enums/account-type"
import "github.com/SolarLabRU/fastpay-go-commons/enums/identity-type"
import "github.com/SolarLabRU/fastpay-go-commons/enums/juridical-type"

type Account struct {
	Address                    string                       `json:"address"`
	State                      state.State                  `json:"state"`
	CurrencyCode               int                          `json:"currencyCode"`
	JuridicalType              juridical_type.JuridicalType `json:"juridicalType"`
	BikBankSetterJuridicalType string                       `json:"bikBankSetterJuridicalType"`
	IdentityType               identity_type.IdentityType   `json:"identityType"`
	Owner                      string                       `json:"owner"`
	Type                       account_type.AccountType     `json:"type"`
	Identifiers                []string                     `json:"identifiers"`
	Encrypted                  bool                         `json:"encrypted"`
	Created                    int64                        `json:"created"`
	PublicKey                  string                       `json:"publicKey"`
	DocType                    string                       `json:"docType"`
}

type Bank struct {
	Address     string      `json:"address"`
	Name        string      `json:"name"`
	BIK         string      `json:"bik"`
	State       state.State `json:"state"`
	CreatedBy   string      `json:"createdBy"`
	IsOwner     bool        `json:"isOwner"`
	Encrypted   bool        `json:"encrypted"`
	IsRegulator bool        `json:"isRegulator"`
	MSPId       string      `json:"mspId"`
	Conf        string      `json:"conf"`
	DocType     string      `json:"docType"`
}
