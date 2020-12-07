package requests

import "fmt"

type CreateClientBankRequest struct {
	TechnicalSignRequest
	Address         string            `yaml:"address" json:"address" valid:"required~ErrorBankAddressNotPassed"`
	BankDisplayName string            `yaml:"bankDisplayName" json:"bankDisplayName" valid:"required~ErrorBankDisplayNameNotPassed"`
	CountryCode     string            `yaml:"countryCode" json:"countryCode"`
	Params          map[string]string `yaml:"params" json:"params"`
}

func (createClientBankRequest *CreateClientBankRequest) SetDefaults() {
	if createClientBankRequest.Params == nil {
		createClientBankRequest.Params = make(map[string]string)
	}
}

func (request *CreateClientBankRequest) String() string {
	return fmt.Sprintf("create%s%s%s", request.Address, request.BankDisplayName, request.CountryCode)
}
