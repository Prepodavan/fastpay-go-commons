package serialising

import (
	"reflect"
)

// Вызывает SetDefaults() у апстрима, если такой метод у апстрима реализован
type DefaultsSetter struct{}

func (ds *DefaultsSetter) ToString(entry *ToStringEntry) (string, error) {
	return entry.Upstream, nil
}

func (ds *DefaultsSetter) FromString(entry *FromStringEntry) (*reflect.Value, error) {
	if converted, ok := entry.Upstream.Interface().(interface{ SetDefaults() }); ok {
		converted.SetDefaults()
	}
	return entry.Upstream, nil
}
