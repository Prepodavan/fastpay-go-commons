package models

// +k8s:deepcopy-gen=true
// +gen * set ring linkedlist slice:"Any,All,Where,Count,Select[CurrencyContractRoutingItem]"
type CurrencyContractRoutingItem struct {
	CurrencyExchangeContract
	AmountInput  int64 `json:"amountInput"`
	AmountOutput int64 `json:"amountOutput"`
}
