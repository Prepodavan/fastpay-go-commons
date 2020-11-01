package serialising

import (
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"reflect"
)

type ToStringInput struct {
	Value      *reflect.Value
	Type       reflect.Type
	Returns    *metadata.ReturnMetadata
	Components *metadata.ComponentMetadata
}

type ToStringEntry struct {
	Input    *ToStringInput
	Upstream string
}

type FromStringInput struct {
	Data       string
	Type       reflect.Type
	Params     *metadata.ParameterMetadata
	Components *metadata.ComponentMetadata
}

type FromStringEntry struct {
	Input    *FromStringInput
	Upstream *reflect.Value
}
