package middlewares

import (
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	roleEnum "github.com/SolarLabRU/fastpay-go-commons/enums/access-role-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
	"github.com/SolarLabRU/fastpay-go-commons/txctx"
)

func (bc *BaseContract) MustHaveMSPID(ctx txctx.ITransactionContext) error {
	if _, err := ctx.GetClientIdentity().GetMSPID(); err != nil {
		return ccErrors.NewBaseError(ccErrors.ErrorGetMspId, ccErrors.WithData(err.Error()))
	}
	return nil
}

func (bc *BaseContract) MustHaveSenderAddress(ctx txctx.ITransactionContext) error {
	_, found, err := ctx.GetClientIdentity().GetAttributeValue("address")
	baseErr := ccErrors.NewBaseError(
		0, //TODO
	)
	switch {
	case err != nil:
		baseErr.Data = err.Error()
	case !found:
		baseErr.Data = "not found"
	}
	return baseErr
}

func (bc *BaseContract) SenderAccess(role roleEnum.AccessRole) func(ctx txctx.ITransactionContext) error {
	return func(ctx txctx.ITransactionContext) error {
		if role == roleEnum.Any {
			return nil
		}
		bank := ctx.GetSenderBank()
		roles := roleEnum.Bank
		for _, v := range bank.Roles {
			roles |= v
		}
		if roles&role == 0 {
			return ccErrors.NewBaseError(ccErrors.ErrorForbidden)
		}
		return nil
	}
}

func (bc *BaseContract) SenderAvailable(ctx txctx.ITransactionContext) error {
	if ctx.GetSenderBank().State != state_enum.Available {
		return ccErrors.NewBaseError(
			ccErrors.ErrorBankNotAvailable,
			ccErrors.WithMessage("Банк отправителя не доступен"),
		)
	}
	return nil
}
