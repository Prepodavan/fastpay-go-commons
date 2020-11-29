package dau

import (
	"bytes"
	"encoding/json"
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func FindAll(stub shim.ChaincodeStubInterface, outputSlice interface{}, query string) error {
	resultsIterator, err := stub.GetQueryResult(query)
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
