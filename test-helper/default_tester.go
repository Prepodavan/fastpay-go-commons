package test_helper

import (
	"encoding/json"
	cc_errors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Запускает TestCase.ChainCodeFunc, проверяя, что
// состояние TestCase.Stub было изменено ожидаемым образом.
// Если TestCase.ErrorCode == nil
// проверяет соответствие ответа функции и TestCase.Response.
// Иначе проверяет, что код возвращенной ошибки
// соответствует TestCase.ErrorCode.
// TestCase.Response и первое возвращенное значение функции могут быть структурами или
// указателями на структуры.
// Строковое представление возвращенной ошибки должно быть json'ом
// по аналогии с cc_errors.JsonSerializableErr.
// Паникует, если:
// 1) TestCase.ChainCodeFunc не принимает на вход
// contractapi.TransactionContextInterface первым аргументом и
// TestCase.Request вторым аргументом (при TestCase.Request != nil).
// 2) TestCase.ChainCodeFunc имеет менее двух возвращаемых значений.
// 3) TestCase.ErrorCode != nil, но второе возвращенное значение не яв-ся ошибкой.
func DefaultTests(t *testing.T, cases []*TestCase) {
	t.Helper()
	t.Log("strict test started")
	defer t.Log("strict test finished")
	for i, tc := range cases {
		t.Logf("case number: %d", i)
		ctx := &contractapi.TransactionContext{}
		ctx.SetStub(tc.Stub)
		ctx.SetClientIdentity(tc.ClientIdentity)
		txID := uuid.New().String()
		tc.Stub.MockTransactionStart(txID)
		args := []reflect.Value{reflect.ValueOf(ctx)}
		if tc.Request != nil {
			args = append(args, reflect.ValueOf(tc.Request))
		}
		results := reflect.ValueOf(tc.ChainCodeFunc).Call(args)
		tc.Stub.MockTransactionEnd(txID)
		if tc.ErrorCode != nil {
			assert.False(t, results[1].IsNil())
			if t.Failed() {
				continue
			}
			err := results[1].Interface().(error)
			var actualErr cc_errors.JsonSerializableErr
			_ = json.Unmarshal([]byte(err.Error()), &actualErr)
			assert.Equal(t, *tc.ErrorCode, actualErr.Code)
		} else {
			assert.True(t, results[1].IsNil())
			assert.True(t, tc.Stub.ExpectationsWhereMet())
			actualResponse := results[0]
			res := reflect.ValueOf(tc.Response)
			if res.Kind() == reflect.Ptr {
				res = res.Elem()
			}
			if actualResponse.Kind() == reflect.Ptr {
				actualResponse = actualResponse.Elem()
			}
			assert.Equal(t, res.Interface(), actualResponse.Interface())
		}
	}
}
