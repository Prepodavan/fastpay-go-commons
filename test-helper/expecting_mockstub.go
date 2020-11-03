package test_helper

import (
	"container/list"
	"encoding/json"
)

// Хранит ожидаемые вызовы изменений состояния.
// Имеет версии методов изменения состояния с проверкой и без проверки ожиданий.
// Любая проверка удаляет вызов из списка ожидаемых, если он был ожидаем,
// иначе записывает вызов в список неожиданных.
type ExpectingWritesMockStub struct {
	*MockStub
	Expected   *list.List
	Unexpected *list.List
}

func NewExpectingWritesMockStub(ms *MockStub) *ExpectingWritesMockStub {
	return &ExpectingWritesMockStub{
		MockStub:   ms,
		Expected:   list.New(),
		Unexpected: list.New(),
	}
}

func (stub *ExpectingWritesMockStub) asPayload(jsonSerializableOrBytes interface{}) (bytes []byte) {
	var ok bool
	if bytes, ok = jsonSerializableOrBytes.([]byte); !ok {
		bytes, _ = json.Marshal(jsonSerializableOrBytes)
	}
	return
}

func (stub *ExpectingWritesMockStub) expect(exp comparableExpectation) {
	stub.Expected.PushBack(exp)
}

func (stub *ExpectingWritesMockStub) checkExpectation(exp comparableExpectation) {
	expected := false
	for el := stub.Expected.Front(); el != nil; el = el.Next() {
		currentExpectation, ok := el.Value.(comparableExpectation)
		if ok && exp.Equals(currentExpectation) {
			expected = true
			stub.Expected.Remove(el)
			break
		}
	}
	if !expected {
		stub.Unexpected.PushBack(exp)
	}
}

// Проверяет ожидание вызова и проксирует вызов в MockStub
func (stub *ExpectingWritesMockStub) PutState(key string, payload []byte) error {
	stub.checkExpectation(&putStateExpectation{
		Key:     key,
		Payload: payload,
	})
	return stub.MockStub.PutState(key, payload)
}

// Проверяет ожидание вызова и проксирует вызов в MockStub
func (stub *ExpectingWritesMockStub) PutPrivateData(col, key string, payload []byte) error {
	stub.checkExpectation(&putPvtDataExpectation{
		Key:        key,
		Collection: col,
		Payload:    payload,
	})
	return stub.MockStub.PutPrivateData(col, key, payload)
}

// Проверяет ожидание вызова и проксирует вызов в MockStub
func (stub *ExpectingWritesMockStub) DelState(key string) error {
	stub.checkExpectation(&delStateExpectation{Key: key})
	return stub.MockStub.DelState(key)
}

// Проверяет ожидание вызова и проксирует вызов в MockStub
func (stub *ExpectingWritesMockStub) DelPrivateData(col, key string) error {
	stub.checkExpectation(&delPvtStateExpectation{Key: key, Collection: col})
	return stub.MockStub.DelPrivateData(col, key)
}

// Проверяет ожидание вызова и проксирует вызов в MockStub
func (stub *ExpectingWritesMockStub) SetEvent(name string, payload []byte) error {
	stub.checkExpectation(&setEventExpectation{Name: name, Payload: payload})
	return stub.MockStub.SetEvent(name, payload)
}

// Возвращает true если на момент вызова не имеется ожиданий и не имеется неожиданностей
func (stub *ExpectingWritesMockStub) ExpectationsWhereMet() bool {
	return stub.Expected.Len() == 0 && stub.Unexpected.Len() == 0
}

// Записывает ожидание вызова в список
func (stub *ExpectingWritesMockStub) ExpectPutState(
	key string,
	jsonSerializableOrBytes interface{},
) (out *ExpectingWritesMockStub) {
	out = stub
	stub.expect(&putStateExpectation{
		Key:     key,
		Payload: stub.asPayload(jsonSerializableOrBytes),
	})
	return
}

// Записывает ожидание вызова в список
func (stub *ExpectingWritesMockStub) ExpectPutPrivateData(
	col, key string,
	jsonSerializableOrBytes interface{},
) (out *ExpectingWritesMockStub) {
	out = stub
	stub.expect(&putPvtDataExpectation{
		Collection: col,
		Key:        key,
		Payload:    stub.asPayload(jsonSerializableOrBytes),
	})
	return
}

// Записывает ожидание вызова в список
func (stub *ExpectingWritesMockStub) ExpectSetEvent(
	name string,
	jsonSerializableOrBytes interface{},
) (out *ExpectingWritesMockStub) {
	out = stub
	stub.expect(&setEventExpectation{
		Name:    name,
		Payload: stub.asPayload(jsonSerializableOrBytes),
	})
	return
}

// Записывает ожидание вызова в список
func (stub *ExpectingWritesMockStub) ExpectDelState(key string) (out *ExpectingWritesMockStub) {
	out = stub
	stub.expect(&delStateExpectation{
		Key: key,
	})
	return
}

// Записывает ожидание вызова в список
func (stub *ExpectingWritesMockStub) ExpectDelPrivateData(
	collection, key string,
) (out *ExpectingWritesMockStub) {
	out = stub
	stub.expect(&delPvtStateExpectation{
		Key:        key,
		Collection: collection,
	})
	return
}

// Проксирует вызов в MockStub без проверки ожиданий
func (stub *ExpectingWritesMockStub) Put(
	key string,
	jsonSerializableOrBytes interface{},
) (out *ExpectingWritesMockStub) {
	out = stub
	_ = stub.MockStub.PutState(key, stub.asPayload(jsonSerializableOrBytes))
	return
}

// Проксирует вызов в MockStub без проверки ожиданий
func (stub *ExpectingWritesMockStub) PutPvt(
	collection, key string,
	jsonSerializableOrBytes interface{},
) (out *ExpectingWritesMockStub) {
	out = stub
	_ = stub.MockStub.PutPrivateData(collection, key, stub.asPayload(jsonSerializableOrBytes))
	return
}

// Проксирует вызов в MockStub без проверки ожиданий
func (stub *ExpectingWritesMockStub) WithEvent(
	name string,
	jsonSerializableOrBytes interface{},
) (out *ExpectingWritesMockStub) {
	out = stub
	_ = stub.MockStub.SetEvent(name, stub.asPayload(jsonSerializableOrBytes))
	return
}
