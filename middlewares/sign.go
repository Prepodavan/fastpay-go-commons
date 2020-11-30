package middlewares

import (
	ccErrors "github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/SolarLabRU/fastpay-go-commons/crypto"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
	"github.com/SolarLabRU/fastpay-go-commons/txctx"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (bc *BaseContract) ValidateTechSign(
	_ contractapi.TransactionContextInterface,
	request requests.TechnicalSign,
) error {
	return bc.validateSign(request.GetTechnicalMsgHash(), request.GetTechnicalAddress(), request.GetTechnicalSig())
}

func (bc *BaseContract) ValidateSigned(_ contractapi.TransactionContextInterface, request requests.Signed) error {
	return bc.validateSign(crypto.Hash(request.String()), request.GetTechnicalAddress(), request.GetTechnicalSig())
}

func (bc *BaseContract) TechnicalAddressShouldBeEqualToSenderAddress(
	ctx txctx.ITransactionContext,
	request requests.TechnicalSign,
) (err error) {
	// TODO Убрать проверку если в сертификате не будет указыватся адрес банка отправителя
	if ctx.GetSenderBank().Address != request.GetTechnicalAddress() {
		err = ccErrors.NewBaseError(ccErrors.ErrorAccountTechnicalNotEqlSender)
	}
	return
}

func (bc *BaseContract) validateSign(msgHash, address string, sign requests.SignDto) error {
	// TODO тесты
	if msgHash == "" || sign.R == "" || sign.S == "" || sign.V == 0 {
		return ccErrors.NewBaseError(ccErrors.ErrorValidateDefault, ccErrors.WithMessage("Сигнатура не передана"))
	}
	isSigned, err := crypto.IsSigned(address, msgHash, sign.R, sign.S, sign.V)
	if err != nil || !isSigned {
		return ccErrors.NewBaseError(ccErrors.ErrorSignVerify)
	}
	return nil
}
