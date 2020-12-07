package requests

type TopupRequest struct {
	TechnicalSignRequest
	AddressTo     string `yaml:"addressTo" json:"addressTo" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount        int64  `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed"`
	CurrencyCode  int    `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Payload       string `yaml:"payload" json:"payload"`
	TransactionId string `yaml:"transactionId" json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
}
