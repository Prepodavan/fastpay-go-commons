package dau

import (
	"bytes"
	"encoding/json"
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func FindAll(stub shim.ChaincodeStubInterface, outputSlice interface{}, query string, opts ...OptionFindAll) error {
	resultsIterator, err := options(opts).selector()(stub, query)
	if err != nil {
		return ccErrors.NewBaseError(
			ccErrors.ErrorGetState,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage("Ошибка при получении данных из БД"),
		)
	}
	result := make([][]byte, 0)
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return ccErrors.NewBaseError(
				ccErrors.ErrorGetState,
				ccErrors.WithData(err.Error()),
				ccErrors.WithMessage("Ошибка при получении данных из БД"),
			)
		}
		result = append(result, response.GetValue())
	}
	full := []byte{'['}
	full = append(full, bytes.Join(result, []byte(","))...)
	full = append(full, ']')
	err = json.Unmarshal(full, outputSlice)
	if err != nil {
		return ccErrors.NewBaseError(
			ccErrors.ErrorJsonUnmarshal,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage("Ошибка при десериализации данных"),
		)
	}
	return nil
}

type options []OptionFindAll

func (opts options) selector() selector {
	for _, opt := range opts {
		if opt.selector != nil {
			return opt.selector
		}
	}
	return defaultSelector
}

type OptionFindAll struct {
	selector selector
}

type selector func(stub shim.ChaincodeStubInterface, query string) (shim.StateQueryIteratorInterface, error)

func defaultSelector(stub shim.ChaincodeStubInterface, query string) (shim.StateQueryIteratorInterface, error) {
	return stub.GetQueryResult(query)
}

func WithPagination(pageSize int32, bookmark string, addrOfPtr **peer.QueryResponseMetadata) OptionFindAll {
	return OptionFindAll{selector: func(
		stub shim.ChaincodeStubInterface,
		query string,
	) (shim.StateQueryIteratorInterface, error) {
		it, md, err := stub.GetQueryResultWithPagination(query, pageSize, bookmark)
		if err == nil && md != nil {
			*addrOfPtr = md
		}
		return it, err
	}}
}
