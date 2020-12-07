package models

// +k8s:deepcopy-gen=true
// +gen * set ring linkedlist slice:"Any,All,Where,Count,Select[CurrencyContractRoutingItem]"
type CurrencyContractRoutingItem struct {
	CurrencyExchangeContract
	AmountInput  int64 `yaml:"amountInput" json:"amountInput"`
	AmountOutput int64 `yaml:"amountOutput" json:"amountOutput"`
}
