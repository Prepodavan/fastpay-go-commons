package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/invoice-state-enum"

type UpdateInvoiceStateRequest struct {
	Number       string                          `yaml:"number" json:"number" valid:"required~ErrorNumberNotPassed,maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	Recipient    string                          `yaml:"recipient" json:"recipient" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode int                             `yaml:"currencyCode" json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	State        invoice_state_enum.InvoiceState `yaml:"state" json:"state" valid:"required~ErrorInvoiceStateNotPassed"`
}
