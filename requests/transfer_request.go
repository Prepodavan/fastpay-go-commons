package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/transaction-type-enum"

type TransferRequest struct {
	AddressFrom   string                                `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To            string                                `yaml:"to" json:"to" valid:"required~ErrorAddressNotPassed,validHex40or64~ErrorAddressOrIdentifierNotFolowingRegex"`
	CurrencyCode  int                                   `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Amount        int64                                 `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed"`
	Payload       string                                `yaml:"payload" json:"payload"`
	MsgHash       string                                `yaml:"msgHash" json:"msgHash"`
	Sig           SignDto                               `yaml:"sig" json:"sig"`
	InvoiceNumber string                                `yaml:"invoiceNumber" json:"invoiceNumber" valid:"maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	Exp           int64                                 `yaml:"exp" json:"exp"`
	TransactionId string                                `yaml:"transactionId" json:"transactionId" valid:"uuidv4"`
	TxType        transaction_type_enum.TransactionType `yaml:"txType" json:"txType"`
}

func (transferRequest *TransferRequest) SetDefaults() {
	if transferRequest.TxType == transaction_type_enum.Undefined {
		transferRequest.TxType = transaction_type_enum.Transfer
	}
}
