package requests

type withMethod interface {
	SetMethodName(methodName string)
}

type withCC interface {
	SetCcName(ccName string)
}

type withCurrencyCode interface {
	SetCurrencyCodeOption(currencyCodeOption int)
}

type Option func(interface{})

func WithCurrency(code int) Option {
	return func(req interface{}) {
		if casted, ok := req.(withCurrencyCode); ok {
			casted.SetCurrencyCodeOption(code)
		}
	}
}

func WithMethod(fun string) Option {
	return func(req interface{}) {
		if casted, ok := req.(withMethod); ok {
			casted.SetMethodName(fun)
		}
	}
}

func WithCC(name string) Option {
	return func(req interface{}) {
		if casted, ok := req.(withCC); ok {
			casted.SetCcName(name)
		}
	}
}
