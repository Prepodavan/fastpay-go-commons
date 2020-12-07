package pools

import (
	"github.com/SolarLabRU/fastpay-go-commons/cache"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/invoker"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type AccountsBalancer struct {
	Accounts
	Balances
}

func NewAccountsBalancer(ctx contractapi.TransactionContextInterface) *AccountsBalancer {
	inv := invoker.NewInvoker(ctx.GetStub())
	return &AccountsBalancer{
		Accounts: Accounts{
			getter: &Base{
				cache:   cache.NewCache(),
				invoker: inv,
			},
		},
		Balances: Balances{
			getter: &Base{
				cache:   cache.NewCache(),
				invoker: inv,
			},
		},
	}
}
