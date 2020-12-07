package requests

type InvoicePaymentRequest struct {
	Number       string `yaml:"number" json:"number" valid:"required~ErrorNumberNotPassed,maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	Recipient    string `yaml:"recipient" json:"recipient" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Payer        string `yaml:"payer" json:"payer" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode int    `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Amount       int64  `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed,range(1|9223372036854775807)~ErrorAmountNegative"`
}
