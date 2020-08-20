package requests

type CreateClientBankRequest struct {
	BankDisplayName string            `json:"bankDisplayName" valid:"required~ErrorBankDisplayNameNotPassed"`
	CountryCode     string            `json:"countryCode"`
	Params          map[string]string `json:"params"`
}

func (createClientBankRequest *CreateClientBankRequest) SetDefaults() {
	if createClientBankRequest.Params == nil {
		createClientBankRequest.Params = make(map[string]string)
	}
}
