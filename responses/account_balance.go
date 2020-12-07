package responses

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"sort"
	"strings"
)

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=github.com/SolarLabRU/fastpay-go-commons/remotes/bus.Response
type AccountBalanceResponse struct {
	Data AccountBalanceData `json:"data"`
	BaseResponse
}

func (in *AccountBalanceResponse) Hash() string {
	return in.Data.Hash()
}

// +k8s:deepcopy-gen=true
type AccountBalanceData struct {
	Items []models.AmountOfBank `json:"items"`
	Total int64                 `json:"total"`
}

func (in *AccountBalanceData) Hash() string {
	banksAddr := make([]string, len(in.Items))
	sort.Strings(banksAddr)
	for _, item := range in.Items {
		banksAddr = append(banksAddr, item.Address)
	}
	return strings.Join(banksAddr, "_")
}
