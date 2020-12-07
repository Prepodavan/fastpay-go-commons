// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Customer

package models

// CustomerSlice is a slice of type Customer. Use it where you would use []Customer.
type CustomerSlice []Customer

// Any verifies that one or more elements of CustomerSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv CustomerSlice) Any(fn func(Customer) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}

// All verifies that all elements of CustomerSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#All
func (rcv CustomerSlice) All(fn func(Customer) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Where returns a new CustomerSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv CustomerSlice) Where(fn func(Customer) bool) (result CustomerSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Count gives the number elements of CustomerSlice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv CustomerSlice) Count(fn func(Customer) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}