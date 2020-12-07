package models

// +gen set ring linkedlist slice:"Any,All,Where,Count"
type Customer struct {
	Identifier          string `json:"identifier"`
	BankAddress         string `json:"bankAddress"`
	BankDisplayName     string `json:"bankDisplayName"`
	CountryCode         string `json:"countryCode"`
	CustomerDisplayName string `json:"customerDisplayName"`
	DocType             string `json:"docType"`
}
