package bus

import (
	"github.com/SolarLabRU/fastpay-go-commons/object"
)

type Response interface {
	object.Hashable
	DeepCopyResponse() Response
}
