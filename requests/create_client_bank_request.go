package requests

import "fmt"

type CreateClientBankRequest struct {
	TechnicalSignRequest
	Address         string            `json:"address" valid:"required~ErrorBankAddressNotPassed"`
	BankDisplayName string            `json:"bankDisplayName" valid:"required~ErrorBankDisplayNameNotPassed"`
	CountryCode     string            `json:"countryCode"`
	Params          map[string]string `json:"params"`
}

func (createClientBankRequest *CreateClientBankRequest) SetDefaults() {
	if createClientBankRequest.Params == nil {
		createClientBankRequest.Params = make(map[string]string)
	}
}

func (request *CreateClientBankRequest) String() string {
	return fmt.Sprintf("create%s%s%s", request.Address, request.BankDisplayName, request.CountryCode)
}
