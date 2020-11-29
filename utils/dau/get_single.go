package dau

import (
	"encoding/json"
	"fmt"
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func Get(stub shim.ChaincodeStubInterface, output interface{}, address string) (found bool, err error) {
	payload, err := stub.GetState(address)
	if err != nil {
		return false, ccErrors.NewBaseError(
			ccErrors.ErrorGetState,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage(fmt.Sprintf("Ошибка при получении данных по адресу %s из БД", address)),
		)
	}
	if len(payload) == 0 {
		return false, nil
	}
	err = json.Unmarshal(payload, output)
	if err != nil {
		return false, ccErrors.NewBaseError(
			ccErrors.ErrorJsonUnmarshal,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage(fmt.Sprintf(
				"Ошибка при десериализации данных (в едином экз) по адресу %s",
				address,
			)),
		)
	}
	return true, nil
}

func Find(stub shim.ChaincodeStubInterface, output interface{}, query string) (found bool, err error) {
	resultsIterator, err := stub.GetQueryResult(query)
	if err != nil {
		return false, ccErrors.NewBaseError(
			ccErrors.ErrorGetState,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage("Ошибка при получении данных (в едином экз) из БД"),
		)
	}
	if !resultsIterator.HasNext() {
		return false, nil
	}
	queryResponse, err := resultsIterator.Next()
	if err != nil {
		return false, ccErrors.NewBaseError(
			ccErrors.ErrorGetState,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage("Ошибка получения данных (в едином экз) "+
				"итератора из ответа по запросу к БД"),
		)
	}
	err = json.Unmarshal(queryResponse.GetValue(), output)
	if err != nil {
		return false, ccErrors.NewBaseError(
			ccErrors.ErrorJsonUnmarshal,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage("Ошибка при десериализации данных (в едином экз)"),
		)
	}
	return true, nil
}
