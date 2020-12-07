package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ExecuteContractRequest struct {
	ContractInfo  models.CurrencyContractRoutingItem `yaml:"contractInfo" json:"contractInfo" valid:"required"`
	AddressFrom   string                             `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To            string                             `yaml:"to" json:"to" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
	Amount        int64                              `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed"`
	TransactionId string                             `yaml:"transactionId" json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
	InvoiceNumber string                             `yaml:"invoiceNumber" json:"invoiceNumber" valid:"maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	Payload       string                             `yaml:"payload" json:"payload"`
}
