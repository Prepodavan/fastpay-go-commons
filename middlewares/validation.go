package middlewares

import (
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/SolarLabRU/fastpay-go-commons/txctx"
	"github.com/SolarLabRU/fastpay-go-commons/validation"
	"github.com/asaskevich/govalidator"
)

func (bc *BaseContract) Validate(ctx txctx.ITransactionContext, request interface{}) error {
	isValid, err := validation.ValidateStruct(request)
	if isValid && err == nil {
		ctx.SetRequest(request)
		return nil
	}
	switch validationErr := err.(type) {
	case govalidator.Errors:
		errName := validationErr[0].Error()
		return ccErrors.NewBaseError(
			ccErrors.ErrorStringCodeMap[errName],
			ccErrors.WithData(err.Error()),
		)
	default:
		return ccErrors.NewBaseError(
			ccErrors.ErrorValidateDefault,
			ccErrors.WithMessage("Ошибка валидации"),
			ccErrors.WithData(err.Error()),
		)
	}
}
