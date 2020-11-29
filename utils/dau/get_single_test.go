package dau

import (
	"github.com/AlekSi/pointer"
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	testHelper "github.com/SolarLabRU/fastpay-go-commons/test-helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	t.Parallel()
	type model struct {
		Msg   string
		Value int8
	}
	type expected struct {
		notFound bool
		result   model
		errCode  *int
	}
	type testCase struct {
		stub     *testHelper.ExpectingWritesMockStub
		expected expected
		selector string
	}
	cases := []testCase{
		{
			selector: H{"selector": H{"Value": 1}}.String(),
			expected: expected{
				result: model{
					Value: 1,
					Msg:   "first",
				},
			},
			stub: testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)).
				Put("first", model{Value: 1, Msg: "first"}).
				Put("second", model{Value: 1, Msg: "second"}).
				Put("third", model{Value: 3, Msg: "third"}),
		},
		{
			selector: H{"selector": H{"Value": 500}}.String(),
			expected: expected{notFound: true},
			stub: testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)).
				Put("first", model{Value: 1, Msg: "first"}).
				Put("second", model{Value: 1, Msg: "second"}).
				Put("third", model{Value: 3, Msg: "third"}),
		},
		{
			selector: H{"Value": 1}.String(),
			expected: expected{errCode: pointer.ToInt(ccErrors.ErrorGetState)},
			stub:     testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)),
		},
	}
	buildTest := func(tc testCase) func(tt *testing.T) {
		return func(tt *testing.T) {
			tt.Parallel()
			var result model
			found, err := Find(tc.stub, &result, tc.selector)
			if tc.expected.errCode != nil {
				assert.NotNil(t, err)
				if t.Failed() {
					return
				}
				assert.Equal(t, err.(*ccErrors.BaseError).Code, *tc.expected.errCode)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expected.notFound, !found)
				if t.Failed() {
					return
				}
				assert.Equal(t, tc.expected.result, result)
			}
		}
	}
	for i, tc := range cases {
		t.Run(testHelper.NumberedTest(i, t.Name()), buildTest(tc))
	}
}

func TestGet(t *testing.T) {
	t.Parallel()
	type model struct {
		Msg   string
		Value int8
	}
	type expected struct {
		notFound bool
		result   model
	}
	type testCase struct {
		stub     *testHelper.ExpectingWritesMockStub
		expected expected
		address  string
	}
	cases := []testCase{
		{
			address: "third",
			expected: expected{
				result: model{
					Value: 3,
					Msg:   "third",
				},
			},
			stub: testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)).
				Put("first", model{Value: 1, Msg: "first"}).
				Put("second", model{Value: 1, Msg: "second"}).
				Put("third", model{Value: 3, Msg: "third"}),
		},
		{
			address: "first",
			expected: expected{
				result: model{
					Value: 1,
					Msg:   "first",
				},
			},
			stub: testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)).
				Put("first", model{Value: 1, Msg: "first"}).
				Put("second", model{Value: 2, Msg: "second"}).
				Put("third", model{Value: 3, Msg: "third"}),
		},
		{
			address:  "0",
			expected: expected{notFound: true},
			stub: testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)).
				Put("first", model{Value: 1, Msg: "first"}).
				Put("second", model{Value: 1, Msg: "second"}).
				Put("third", model{Value: 3, Msg: "third"}),
		},
	}
	buildTest := func(tc testCase) func(tt *testing.T) {
		return func(tt *testing.T) {
			tt.Parallel()
			var result model
			found, err := Get(tc.stub, &result, tc.address)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected.notFound, !found)
			if t.Failed() {
				return
			}
			assert.Equal(t, tc.expected.result, result)
		}
	}
	for i, tc := range cases {
		t.Run(testHelper.NumberedTest(i, t.Name()), buildTest(tc))
	}
}
