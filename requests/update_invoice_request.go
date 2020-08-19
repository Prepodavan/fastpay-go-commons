package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/invoice-state-enum"

type UpdateInvoiceRequest struct {
	Number       string                          `json:"number" valid:"required~60352,maxstringlength(255)~60353"`
	Recipient    string                          `json:"recipient" valid:"required~60302,validHex40~60301"`
	Payer        string                          `json:"payer" valid:"required~60302,validHex40~60301"`
	State        invoice_state_enum.InvoiceState `json:"state" valid:"required~60355"`
	CurrencyCode int                             `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	ErrorCode    int                             `json:"errorCode"`
}
