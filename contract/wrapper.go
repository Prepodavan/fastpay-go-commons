package contract

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strings"
)

// Проксирует вызовы в CC. Если результатом Invoke была ошибка, строковое представлением
// которой яв-ся json, то итоговый peer.Response будет содержать сообщение с
// вырезанным json из текста оригинальной ошибки.
// Иначе будет возвращен исходный ответ
type JsonErrWrapper struct {
	CC *contractapi.ContractChaincode
}

func (w *JsonErrWrapper) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return w.CC.Init(stub)
}

func (w *JsonErrWrapper) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	res := w.CC.Invoke(stub)
	if res.Status != shim.ERROR {
		return res
	}
	start := strings.Index(res.Message, "{")
	end := strings.LastIndex(res.Message, "}")
	if start == -1 || end == -1 {
		return res
	}
	return peer.Response{
		Status:  res.Status,
		Message: res.Message[start : end+1],
		Payload: res.Payload,
	}
}
