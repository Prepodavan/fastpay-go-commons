package pools

import (
	"encoding/json"
	"errors"
	"github.com/SolarLabRU/fastpay-go-commons/object"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/bus"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/invoker"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/rcc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var errUnexpectedRequest = errors.New("unexpected request")

type cacheImpl map[string]object.Hashable

func (c cacheImpl) Has(request object.Hashable) (ok bool) {
	_, ok = c[request.Hash()]
	return
}

func (c cacheImpl) Get(request object.Hashable) (response object.Hashable) {
	response, _ = c[request.Hash()]
	return
}

func (c cacheImpl) Add(request, response object.Hashable) {
	c[request.Hash()] = response
}

type invkr struct {
	m map[string]interface{}
}

func (i *invkr) Respond(req invoker.Request, res bus.Response) *invkr {
	i.m[req.Handler().Func()+req.Handler().Name()] = res
	return i
}

func (i *invkr) Err(req invoker.Request, err error) *invkr {
	i.m[req.Handler().Func()+req.Handler().Name()] = err
	return i
}

func (i *invkr) Request(response interface{}, request invoker.Request) error {
	res, ok := i.m[request.Handler().Func()+request.Handler().Name()]
	if !ok {
		return errUnexpectedRequest
	}
	if err, ok := res.(error); ok {
		return err
	}
	d, err := json.Marshal(res)
	if err != nil {
		return err
	}
	err = json.Unmarshal(d, response)
	return err
}

func TestBase_Get(t *testing.T) {
	inv := &invkr{m: make(map[string]interface{})}
	cch := make(cacheImpl)
	b := &Base{
		cache:   cch,
		invoker: inv,
	}
	_, err := b.Get(nil, &Payload{})
	assert.Equal(t, errUnexpectedRequest, err)
	if t.Failed() {
		return
	}
	inv.Err(&Payload{Data: []string{"1"}}, errors.New("1"))
	_, err = b.Get(nil, &Payload{Data: []string{"1"}})
	assert.NotNil(t, err)
	if t.Failed() {
		return
	}
	assert.Equal(t, "1", err.Error())
	if t.Failed() {
		return
	}
	expectedRes := &Payload{Data: []string{"2"}}
	inv.Respond(&Payload{Data: []string{"req2"}}, expectedRes)
	actual, err := b.Get(&Payload{}, &Payload{Data: []string{"req2"}})
	assert.Nil(t, err)
	if t.Failed() {
		return
	}
	assert.Equal(t, expectedRes, actual)
	if t.Failed() {
		return
	}
	actual.(*Payload).Data[0] = "3"
	assert.Equal(t, expectedRes.Data[0], "2")
	assert.Equal(t, actual.(*Payload).Data[0], "3")
	if t.Failed() {
		return
	}
	actual, err = b.Get(&Payload{}, &Payload{Data: []string{"req2"}})
	assert.Nil(t, err)
	if t.Failed() {
		return
	}
	assert.Equal(t, expectedRes, actual)
}

type Payload struct {
	Data []string
}

func (p *Payload) DeepCopyResponse() bus.Response {
	d := make([]string, len(p.Data))
	copy(d, p.Data)
	return &Payload{Data: d}
}

func (p *Payload) Hash() string {
	return strings.Join(p.Data, "")
}

func (p *Payload) Handler() rcc.Handler {
	return &handler{
		name: p.Hash(),
		fun:  p.Hash(),
	}
}

type handler struct {
	name string
	fun  string
}

func (h *handler) Name() string {
	return h.name
}

func (h *handler) Func() string {
	return h.fun
}
