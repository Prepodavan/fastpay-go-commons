package requests

type WithdrawConfirmRequest struct {
	TechnicalSignRequest
	AddressFrom   string `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	TxId          string `yaml:"txId" json:"txId" valid:"required~ErrorTxIdSNotPassed"`
	CurrencyCode  int    `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	TransactionId string `yaml:"transactionId" json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
}

type WithdrawRejectRequest WithdrawConfirmRequest
