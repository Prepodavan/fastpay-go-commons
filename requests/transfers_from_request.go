package requests

type TransfersFromRequest struct {
	AddressFrom   string            `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To            []TransfersFromTo `yaml:"to" json:"to" valid:"required"`
	CurrencyCode  int               `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	TransactionId string            `yaml:"transactionId" json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
}

type TransfersFromTo struct {
	To      string `yaml:"to" json:"to" valid:"required~ErrorAddressNotPassed,validHex40or64~ErrorAddressOrIdentifierNotFolowingRegex"`
	Amount  int64  `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed"`
	Payload string `yaml:"payload" json:"payload"`
}
