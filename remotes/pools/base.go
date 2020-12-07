package pools

import (
	"github.com/SolarLabRU/fastpay-go-commons/object"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/bus"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/invoker"
)

type Invoker interface {
	Request(response interface{}, request invoker.Request) error
}

type Request interface {
	object.Hashable
	invoker.Request
}

type Cache interface {
	Add(request, response object.Hashable)
	Has(request object.Hashable) bool
	Get(request object.Hashable) (response object.Hashable)
}

type Base struct {
	cache   Cache
	invoker Invoker
}

func (b *Base) Get(responseType bus.Response, request Request) (found bus.Response, err error) {
	if b.cache != nil && b.cache.Has(request) {
		found = b.cache.Get(request).(bus.Response).DeepCopyResponse()
		return
	}
	err = b.invoker.Request(responseType, request)
	if err == nil && b.cache != nil {
		b.cache.Add(request, responseType.DeepCopyResponse())
		found = responseType
	}
	return
}
