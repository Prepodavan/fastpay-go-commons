package pools

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/bus"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
	"github.com/SolarLabRU/fastpay-go-commons/responses"
)

type Accounts struct {
	getter *Base
}

func (acc *Accounts) Account(addr string) (account *models.Account, err error) {
	var response bus.Response
	response, err = acc.getter.Get(&responses.AccountResponse{}, &requests.GetByAddressRequest{Address: addr})
	if err != nil {
		return
	}
	account = &response.(*responses.AccountResponse).Data
	return
}
