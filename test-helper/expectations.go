package test_helper

type comparableExpectation interface {
	Equals(other interface{}) bool
}

type putStateExpectation struct {
	Key     string
	Payload []byte
}

func (ex *putStateExpectation) Equals(other interface{}) (eq bool) {
	switch casted := other.(type) {
	case putStateExpectation:
		eq = casted.Key == ex.Key && string(casted.Payload) == string(ex.Payload)
	case *putStateExpectation:
		eq = casted.Key == ex.Key && string(casted.Payload) == string(ex.Payload)
	}
	return
}

type putPvtDataExpectation struct {
	Collection string
	Key        string
	Payload    []byte
}

func (ex *putPvtDataExpectation) Equals(other interface{}) (eq bool) {
	switch casted := other.(type) {
	case putPvtDataExpectation:
		eq = casted.Collection == ex.Collection && casted.Key == ex.Key && string(casted.Payload) == string(ex.Payload)
	case *putPvtDataExpectation:
		eq = casted.Collection == ex.Collection && casted.Key == ex.Key && string(casted.Payload) == string(ex.Payload)
	}
	return
}

type setEventExpectation struct {
	Name    string
	Payload []byte
}

func (ex *setEventExpectation) Equals(other interface{}) (eq bool) {
	switch casted := other.(type) {
	case setEventExpectation:
		eq = casted.Name == ex.Name && string(casted.Payload) == string(ex.Payload)
	case *setEventExpectation:
		eq = casted.Name == ex.Name && string(casted.Payload) == string(ex.Payload)
	}
	return
}

type delStateExpectation struct {
	Key string
}

func (ex *delStateExpectation) Equals(other interface{}) (eq bool) {
	switch casted := other.(type) {
	case delStateExpectation:
		eq = casted.Key == ex.Key
	case *delStateExpectation:
		eq = casted.Key == ex.Key
	}
	return
}

type delPvtStateExpectation struct {
	Key        string
	Collection string
}

func (ex *delPvtStateExpectation) Equals(other interface{}) (eq bool) {
	switch casted := other.(type) {
	case delPvtStateExpectation:
		eq = casted.Collection == ex.Collection && casted.Key == ex.Key
	case *delPvtStateExpectation:
		eq = casted.Collection == ex.Collection && casted.Key == ex.Key
	}
	return
}
