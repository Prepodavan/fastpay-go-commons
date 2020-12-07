package models

// +k8s:deepcopy-gen=true
type AmountOfBank struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}
