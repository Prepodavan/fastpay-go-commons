package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/remotes/ccs"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/rcc"
)

type GetByAddressRequest struct {
	Address string `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}

func (g *GetByAddressRequest) Hash() string {
	return g.Address
}

func (g *GetByAddressRequest) Handler() rcc.Handler {
	const fun = "getByAddress"
	return &handler{
		fun:  fun,
		name: ccs.Accounts,
	}
}
