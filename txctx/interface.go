package txctx

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ITransactionContext interface {
	contractapi.TransactionContextInterface
	GetSenderBank() *models.Bank
	SetSenderBank(senderBank *models.Bank)
	GetRequest() interface{}
	SetRequest(request interface{})
	GetValue(key string) (v interface{}, ok bool)
	SetValue(key string, value interface{})
}
