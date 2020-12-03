package middlewares

import (
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/SolarLabRU/fastpay-go-commons/txctx"
	"time"
)

func (bc *BaseContract) MustHaveTimestamp(ctx txctx.ITransactionContext) error {
	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return ccErrors.NewBaseError(
			ccErrors.ErrorGetTxTime,
			ccErrors.WithData(err.Error()),
			ccErrors.WithMessage("Ошибка при получении времени создания транзакции"),
		)
	}
	ctx.SetTimestamp(time.Unix(0, timestamp.GetSeconds()*1000+int64(timestamp.GetNanos()/1e6)))
	return nil
}
