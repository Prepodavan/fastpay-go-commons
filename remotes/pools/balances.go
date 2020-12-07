package pools

import (
	"github.com/SolarLabRU/fastpay-go-commons/remotes/bus"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
	"github.com/SolarLabRU/fastpay-go-commons/responses"
)

type Balances struct {
	getter *Base
}

func (bal *Balances) Balance(addr string, currency int) (account *responses.AccountBalanceData, err error) {
	var response bus.Response
	response, err = bal.getter.Get(&responses.AccountBalanceResponse{}, &requests.GetAccountBalanceRequest{
		Address:      addr,
		CurrencyCode: currency,
	})
	if err != nil {
		return
	}
	account = &response.(*responses.AccountBalanceResponse).Data
	return
}
