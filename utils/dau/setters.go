package dau

import (
	"encoding/json"
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type Addressable interface {
	GetAddress() string
}

func Set(stub shim.ChaincodeStubInterface, input Addressable) error {
	return SetAny(stub, input, input.GetAddress())
}

func SetAny(stub shim.ChaincodeStubInterface, input interface{}, address string) error {
	payload, err := json.Marshal(input)
	if err != nil {
		return ccErrors.NewBaseError(ccErrors.ErrorJsonMarshal, ccErrors.WithData(err.Error()))
	}
	err = stub.PutState(address, payload)
	if err != nil {
		return ccErrors.NewBaseError(ccErrors.ErrorPutState, ccErrors.WithData(err.Error()))
	}
	return nil
}
