package responses

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=github.com/SolarLabRU/fastpay-go-commons/remotes/bus.Response
type AccountResponse struct {
	Data models.Account `yaml:"data" json:"data"`
	BaseResponse
}

func (in *AccountResponse) Hash() string {
	return in.Data.Hash()
}
