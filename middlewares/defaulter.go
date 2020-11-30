package middlewares

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (bc *BaseContract) SetDefaults(_ contractapi.TransactionContextInterface, request interface{}) {
	if settable, ok := request.(interface{ SetDefaults() }); ok {
		settable.SetDefaults()
	}
	return
}
