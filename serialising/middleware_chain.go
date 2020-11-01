package serialising

import (
	"container/list"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/hyperledger/fabric-contract-api-go/serializer"
	"reflect"
)

// Middleware описывает сериалайзер, который может быть звеном цепочки.
// Принимает на вход выхлоп предыдущего звена и общие входные данные сериалайзера.
type Middleware interface {
	ToString(entry *ToStringEntry) (string, error)
	FromString(entry *FromStringEntry) (*reflect.Value, error)
}

// Цепочка сериалайзеров. Каждый не стартовый сериалайзер получает на вход
// выхлоп предыдущего сериалайзера.
// Если middleware в цепочке возвращает ошибку выполнение останавливается.
type MiddlewaresChain struct {
	start serializer.TransactionSerializer
	chain *list.List
}

func StartChain(starter serializer.TransactionSerializer) *MiddlewaresChain {
	return &MiddlewaresChain{
		start: starter,
		chain: list.New(),
	}
}

// Добавляет в цепочку mc данный middleware.
func (mc *MiddlewaresChain) Next(middleware Middleware) *MiddlewaresChain {
	mc.chain.PushBack(middleware)
	return mc
}

func (mc *MiddlewaresChain) FromString(
	data string,
	typ reflect.Type,
	params *metadata.ParameterMetadata,
	components *metadata.ComponentMetadata,
) (out reflect.Value, err error) {
	out, err = mc.start.FromString(data, typ, params, components)
	if err != nil {
		return
	}
	input := &FromStringInput{
		Data:       data,
		Type:       typ,
		Params:     params,
		Components: components,
	}
	for current := mc.chain.Front(); current != nil; current = current.Next() {
		ser := current.Value.(Middleware)
		val, err := ser.FromString(&FromStringEntry{
			Input:    input,
			Upstream: &out,
		})
		if err != nil {
			return reflect.Value{}, err
		}
		out = *val
	}
	return
}

func (mc *MiddlewaresChain) ToString(
	val reflect.Value,
	typ reflect.Type,
	returns *metadata.ReturnMetadata,
	components *metadata.ComponentMetadata,
) (out string, err error) {
	out, err = mc.start.ToString(val, typ, returns, components)
	if err != nil {
		return
	}
	input := &ToStringInput{
		Value:      &val,
		Type:       typ,
		Returns:    returns,
		Components: components,
	}
	for current := mc.chain.Front(); current != nil; current = current.Next() {
		ser := current.Value.(Middleware)
		out, err = ser.ToString(&ToStringEntry{
			Input:    input,
			Upstream: out,
		})
		if err != nil {
			return
		}
	}
	return
}
