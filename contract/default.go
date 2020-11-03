package contract

import (
	"github.com/SolarLabRU/fastpay-go-commons/serialising"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/serializer"
)

// Создает чейнкод с json схемой данных.
// Добавляет мидлвары валидации serialising.Validator и
// установщика дефолтных значений в запрос serialising.DefaultsSetter.
// Оборачивает полученный чейнкод в JsonErrWrapper.
func Default(c contractapi.ContractInterface) (shim.Chaincode, error) {
	cc, err := contractapi.NewChaincode(c)
	if err != nil {
		return nil, err
	}
	cc.TransactionSerializer = serialising.StartChain(&serializer.JSONSerializer{}).
		Next(&serialising.Validator{}).
		Next(&serialising.DefaultsSetter{})
	return &JsonErrWrapper{CC: cc}, nil
}
