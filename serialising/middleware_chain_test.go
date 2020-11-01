package serialising

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/serializer"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)

func TestMiddlewaresChain_FromString(t *testing.T) {
	chain := StartChain(&serializer.JSONSerializer{}).
		Next(&positiveChecker{})
	chainWithStringer := StartChain(chain).Next(&intToStringer{})
	cases := []struct {
		input              *FromStringInput
		chain              *MiddlewaresChain
		expectedErr        error
		expectedValid      bool
		expectedSuccessful inputData
	}{
		{
			input: &FromStringInput{
				Data: `{"Int": -1, "String": "a"}`,
				Type: reflect.TypeOf(&inputData{}),
			},
			chain:         chain,
			expectedErr:   fmt.Errorf("not positive: %d", -1),
			expectedValid: false,
		},
		{
			input: &FromStringInput{
				Data: `{"Int": 100, "String": "a"}`,
				Type: reflect.TypeOf(&inputData{}),
			},
			chain:         chain,
			expectedErr:   nil,
			expectedValid: true,
			expectedSuccessful: inputData{
				Int:    100,
				String: "a",
			},
		},
		{
			input: &FromStringInput{
				Data: `{"Int": 100, "String": "a"}`,
				Type: reflect.TypeOf(&inputData{}),
			},
			chain:         chainWithStringer,
			expectedErr:   nil,
			expectedValid: true,
			expectedSuccessful: inputData{
				Int:    100,
				String: "100",
			},
		},
	}
	var i int
	var ts struct {
		input              *FromStringInput
		chain              *MiddlewaresChain
		expectedErr        error
		expectedValid      bool
		expectedSuccessful inputData
	}
	for i, ts = range cases {
		out, err := ts.chain.FromString(ts.input.Data, ts.input.Type, ts.input.Params, ts.input.Components)
		if ts.expectedErr != nil {
			assert.NotNil(t, err)
			if t.Failed() {
				break
			}
			assert.Equal(t, ts.expectedErr.Error(), err.Error())
			if t.Failed() {
				break
			}
		}
		if ts.expectedValid {
			assert.NotNil(t, out)
			if t.Failed() {
				break
			}
		}
		assert.Equal(t, ts.expectedValid, out.IsValid())
		if t.Failed() {
			break
		}
		if !ts.expectedValid {
			continue
		}
		assert.Equal(t, ts.expectedSuccessful, *(out.Interface().(*inputData)))
		if t.Failed() {
			break
		}
	}
	if t.Failed() {
		t.Logf("case number: %d", i)
	}
}

func TestMiddlewaresChain_ToString(t *testing.T) {
	m := map[string]interface{}{
		"1":    "hi",
		"blah": "other",
	}
	expected := `<html>{
	"1": "hi",
	"blah": "other"
}</html>`
	chain := StartChain(&serializer.JSONSerializer{}).
		Next(&indenting{}).
		Next(&htmlDecorator{})
	val := reflect.ValueOf(m)
	actual, err := chain.ToString(val, val.Type(), nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

type inputData struct {
	Int    int
	String string
}

type positiveChecker struct {
	skippingMiddleware
}

func (*positiveChecker) FromString(entry *FromStringEntry) (*reflect.Value, error) {
	if data := entry.Upstream.Interface().(*inputData); data.Int <= 0 {
		return &reflect.Value{}, fmt.Errorf("not positive: %d", data.Int)
	}
	return entry.Upstream, nil
}

type intToStringer struct {
	skippingMiddleware
}

func (*intToStringer) FromString(entry *FromStringEntry) (*reflect.Value, error) {
	data := entry.Upstream.Interface().(*inputData)
	data.String = strconv.Itoa(data.Int)
	return entry.Upstream, nil
}

type indenting struct {
	skippingMiddleware
}

func (*indenting) ToString(entry *ToStringEntry) (out string, err error) {
	var m map[string]interface{}
	err = json.Unmarshal([]byte(entry.Upstream), &m)
	if err != nil {
		return
	}
	var bytes []byte
	bytes, err = json.MarshalIndent(m, "", "\t")
	out = string(bytes)
	return
}

type htmlDecorator struct {
	skippingMiddleware
}

func (*htmlDecorator) ToString(entry *ToStringEntry) (string, error) {
	return "<html>" + entry.Upstream + "</html>", nil
}

type skippingMiddleware struct{}

func (*skippingMiddleware) ToString(entry *ToStringEntry) (string, error) {
	return entry.Upstream, nil
}

func (*skippingMiddleware) FromString(entry *FromStringEntry) (*reflect.Value, error) {
	return entry.Upstream, nil
}
