package dau

import (
	"github.com/AlekSi/pointer"
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	testHelper "github.com/SolarLabRU/fastpay-go-commons/test-helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAll(t *testing.T) {
	t.Parallel()
	type model struct {
		Msg   string
		Value int8
	}
	type expected struct {
		result  []model
		errCode *int
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
				result: []model{
					{
						Value: 1,
						Msg:   "first",
					},
					{
						Value: 1,
						Msg:   "second",
					},
				},
			},
			stub: testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)).
				Put("first", model{Value: 1, Msg: "first"}).
				Put("second", model{Value: 1, Msg: "second"}).
				Put("third", model{Value: 3, Msg: "third"}),
		},
		{
			selector: H{"selector": H{"Value": 500}}.String(),
			expected: expected{result: make([]model, 0)},
			stub: testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)).
				Put("first", model{Value: 1, Msg: "first"}).
				Put("second", model{Value: 1, Msg: "second"}).
				Put("third", model{Value: 3, Msg: "third"}),
		},
		{
			selector: H{"Value": 1}.String(),
			expected: expected{
				errCode: pointer.ToInt(ccErrors.ErrorGetState),
			},
			stub: testHelper.NewExpectingWritesMockStub(testHelper.NewMockStub(t.Name(), nil)),
		},
	}
	buildTest := func(tc testCase) func(tt *testing.T) {
		return func(tt *testing.T) {
			tt.Parallel()
			var result []model
			err := FindAll(tc.stub, &result, tc.selector)
			if tc.expected.errCode != nil {
				assert.NotNil(t, err)
				if t.Failed() {
					return
				}
				assert.Equal(t, err.(*ccErrors.BaseError).Code, *tc.expected.errCode)
			} else {
				assert.Nil(t, err)
				if t.Failed() {
					return
				}
				assert.Equal(t, len(tc.expected.result), len(result))
				if t.Failed() {
					return
				}
				for j := range tc.expected.result {
					assert.Equalf(t, tc.expected.result[j], result[j], "result number: %d", j)
				}
			}
		}
	}
	for i, tc := range cases {
		t.Run(testHelper.NumberedTest(i, t.Name()), buildTest(tc))
	}
}
