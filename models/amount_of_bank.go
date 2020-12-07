package models

// +k8s:deepcopy-gen=true
type AmountOfBank struct {
	Address string `yaml:"address" json:"address"`
	Amount  int64  `yaml:"amount" json:"amount"`
}
