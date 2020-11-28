package cc_errors

func NewBaseError(code int, opts ...BaseErrorOption) *BaseError {
	opts = append([]BaseErrorOption{
		WithDefaultMessage(),
	}, opts...)
	be := &BaseError{Code: code}
	for _, opt := range opts {
		opt(be)
	}
	return be
}

type BaseErrorOption func(*BaseError)

func WithDefaultMessage() BaseErrorOption {
	return func(be *BaseError) {
		be.Message = ErrorCodeMessagesMap[be.Code]
	}
}

func WithMessage(msg string) BaseErrorOption {
	return func(be *BaseError) {
		be.Message = msg
	}
}

func WithData(data string) BaseErrorOption {
	return func(be *BaseError) {
		be.Data = data
	}
}
