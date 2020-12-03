package remotes

import (
	"github.com/SolarLabRU/fastpay-go-commons/remotes/invoker"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct{}

func (c *Contract) BuildInvoker(ctx contractapi.TransactionContextInterface) *invoker.Invoker {
	return invoker.NewInvoker(ctx.GetStub())
}
