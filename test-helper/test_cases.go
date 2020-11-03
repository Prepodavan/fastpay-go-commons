package test_helper

import (
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type MockStubWithExpectations interface {
	shim.ChaincodeStubInterface
	MockTransactionStart(txID string)
	MockTransactionEnd(txID string)
	ExpectationsWhereMet() bool
}

type TestCase struct {
	Stub                             MockStubWithExpectations
	ClientIdentity                   cid.ClientIdentity
	ErrorCode                        *int
	ChainCodeFunc, Request, Response interface{}
}
