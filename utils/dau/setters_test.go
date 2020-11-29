package dau

import (
	testHelper "github.com/SolarLabRU/fastpay-go-commons/test-helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetAny(t *testing.T) {
	t.Parallel()
	model := struct {
		Msg   string
		Value int8
	}{
		Msg:   "10",
		Value: 10,
	}
	stub := testHelper.
		NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)).
		ExpectPutState("1", model)
	stub.MockTransactionStart("1")
	err := SetAny(stub, model, "1")
	stub.MockTransactionEnd("1")
	assert.Nil(t, err)
	assert.True(t, stub.ExpectationsWhereMet())
}
