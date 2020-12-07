package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/ccs"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/rcc"
)

type GetAccountBalanceRequest struct {
	Address      string `yaml:"address" json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode int    `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}

func (g *GetAccountBalanceRequest) Hash() string {
	return fmt.Sprintf("%d_%s", g.CurrencyCode, g.Address)
}

func (g *GetAccountBalanceRequest) Handler() rcc.Handler {
	return &handler{
		fun:  "getAccountBalance",
		name: fmt.Sprintf("%s%d", ccs.Transactions, g.CurrencyCode),
	}
}
