package txctx

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

type TransactionContext struct {
	contractapi.TransactionContext
	Values     map[string]interface{}
	Request    interface{}
	SenderBank *models.Bank
	Timestamp  time.Time
}

func (ctx *TransactionContext) GetTimestamp() time.Time {
	return ctx.Timestamp
}

func (ctx *TransactionContext) SetTimestamp(timestamp time.Time) {
	ctx.Timestamp = timestamp
}

func (ctx *TransactionContext) GetSenderBank() *models.Bank {
	return ctx.SenderBank
}

func (ctx *TransactionContext) SetSenderBank(senderBank *models.Bank) {
	ctx.SenderBank = senderBank
}

func (ctx *TransactionContext) GetRequest() interface{} {
	return ctx.Request
}

func (ctx *TransactionContext) SetRequest(request interface{}) {
	ctx.Request = request
}

func (ctx *TransactionContext) Value(key string) (v interface{}) {
	if ctx.Values == nil {
		ctx.Values = make(map[string]interface{})
	}
	v, _ = ctx.Values[key]
	return
}

func (ctx *TransactionContext) GetValue(key string) (v interface{}, ok bool) {
	if ctx.Values == nil {
		ctx.Values = make(map[string]interface{})
	}
	v, ok = ctx.Values[key]
	return
}

func (ctx *TransactionContext) SetValue(key string, value interface{}) {
	if ctx.Values == nil {
		ctx.Values = make(map[string]interface{})
	}
	ctx.Values[key] = value
}
