package models

// +gen set ring linkedlist slice:"Any,All,Where,Count"
type Customer struct {
	Identifier          string `yaml:"identifier" json:"identifier"`
	BankAddress         string `yaml:"bankAddress" json:"bankAddress"`
	BankDisplayName     string `yaml:"bankDisplayName" json:"bankDisplayName"`
	CountryCode         string `yaml:"countryCode" json:"countryCode"`
	CustomerDisplayName string `yaml:"customerDisplayName" json:"customerDisplayName"`
	DocType             string `yaml:"docType" json:"docType"`
}
