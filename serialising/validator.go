package serialising

import (
	cc_errors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/SolarLabRU/fastpay-go-commons/validation"
	"github.com/asaskevich/govalidator"
	"reflect"
)

// Валидирует апстрим
type Validator struct{}

func (v *Validator) ToString(entry *ToStringEntry) (string, error) {
	return entry.Upstream, nil
}

func (v *Validator) FromString(entry *FromStringEntry) (*reflect.Value, error) {
	if !entry.Upstream.IsValid() || entry.Upstream.IsNil() {
		return entry.Upstream, nil
	}
	isValid, err := validation.ValidateStruct(entry.Upstream.Interface())
	if isValid && err == nil {
		return entry.Upstream, nil
	}
	switch validationErr := err.(type) {
	case govalidator.Errors:
		errName := validationErr[0].Error()
		return nil, &cc_errors.JsonSerializableErr{
			Code:    cc_errors.ErrorMessages[errName].Code,
			Message: cc_errors.ErrorMessages[errName].Message,
			Details: err,
		}
	default:
		return nil, &cc_errors.JsonSerializableErr{
			Code:    cc_errors.ErrorValidateDefault,
			Message: "Ошибка валидации",
			Details: err,
		}
	}
}
