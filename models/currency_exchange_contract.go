package models

import "math"

// +k8s:deepcopy-gen=true
// +gen * set ring linkedlist slice:"Any,Where,Count,Aggregate[AddressOfAccountSlice],Aggregate[CurrencyContractRoutingItem],All"
type CurrencyExchangeContract struct {
	CurrencyExchangeContractMutable
	BankAddress     string `yaml:"bankAddress" json:"bankAddress"`
	BankDisplayName string `yaml:"bankDisplayName" json:"bankDisplayName"`
	DocType         string `yaml:"docType" json:"docType"`
}

func (c *CurrencyExchangeContract) ListAccountsAddresses() AddressOfAccountSlice {
	return AddressOfAccountSlice{
		AddressOfAccount(c.AddressAccountBuy),
		AddressOfAccount(c.AddressAccountSell),
		AddressOfAccount(c.AddressCommission),
	}
}

func (c *CurrencyExchangeContract) CalcCommission(amountOutput int64) float64 {
	calcCommission := float64(amountOutput) * c.Price * c.FractionalCommission / (1 - c.FractionalCommission)
	if c.MaxCommission == 0 {
		return calcCommission
	}
	return math.Min(calcCommission, float64(c.MaxCommission))
}
