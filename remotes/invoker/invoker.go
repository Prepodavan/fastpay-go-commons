package invoker

import (
	"encoding/json"
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/rcc"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type Invoker struct {
	stub shim.ChaincodeStubInterface
}

type Request interface {
	Handler() rcc.Handler
}

func NewInvoker(stub shim.ChaincodeStubInterface, opts ...Option) (i *Invoker) {
	i = &Invoker{stub: stub}
	for _, opt := range opts {
		opt(i)
	}
	return
}

type Option func(invoker *Invoker)

func (invoker *Invoker) Request(responseAddr interface{}, req Request) error {
	return invoker.Invoke(req.Handler().Name(), req.Handler().Func(), req, responseAddr)
}

func (invoker *Invoker) Invoke(fun, cc string, request, responseAddr interface{}) error {
	req, err := invoker.asPayload(request)
	if err != nil {
		return ccErrors.NewBaseError(
			ccErrors.ErrorDefault,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage("Ошибка парсинга входных параметров"),
		)
	}
	args := [][]byte{
		[]byte(fun),
		req,
	}
	res := invoker.stub.InvokeChaincode(cc, args, "")
	switch res.GetStatus() {
	case shim.ERROR:
		var be ccErrors.BaseError
		err := json.Unmarshal([]byte(res.GetMessage()), &be)
		if err != nil {
			return ccErrors.NewBaseError(
				ccErrors.ErrorDefault,
				ccErrors.WithData(err.Error()),
				ccErrors.WithMessage("Ошибка при вызове чейнкода"),
			)
		} else {
			return &be
		}
	default:
		err = json.Unmarshal(res.GetPayload(), responseAddr)
		if err != nil {
			return ccErrors.NewBaseError(
				ccErrors.ErrorJsonUnmarshal,
				ccErrors.WithData(err.Error()),
				ccErrors.WithMessage("Ошибка десерилизации ответа вовремя вызова чейнкода"),
			)
		}
	}
	return nil
}

func (invoker *Invoker) asPayload(obj interface{}) (b []byte, err error) {
	var ok bool
	if b, ok = obj.([]byte); ok {
		return
	}
	return json.Marshal(obj)
}
