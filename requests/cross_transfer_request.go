package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type CrossTransferRequest struct {
	Routes              []models.CurrencyContractRoutingItem `yaml:"routes" json:"routes" valid:"required"`
	AddressFrom         string                               `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To                  string                               `yaml:"to" json:"to" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
	EncryptedSecretKeys []models.AccountSecretKey            `yaml:"encryptedSecretKeys" json:"encryptedSecretKeys"`
	Amount              int64                                `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed"`
	CurrencyCodeFrom    int                                  `yaml:"currencyCodeFrom" json:"currencyCodeFrom" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CurrencyCodeTo      int                                  `yaml:"currencyCodeTo" json:"currencyCodeTo" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CustomerIdentifier  string                               `yaml:"customerIdentifier" json:"customerIdentifier" valid:"validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode         string                               `yaml:"countryCode" json:"countryCode" valid:"stringlength(3|3)"`
	Payload             string                               `yaml:"payload" json:"payload"`
	InvoiceNumber       string                               `yaml:"invoiceNumber" json:"invoiceNumber" valid:"maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	MsgHash             string                               `yaml:"msgHash" json:"msgHash" valid:"required"`
	Sig                 SignDto                              `yaml:"sig" json:"sig" valid:"required"`
	Exp                 int64                                `yaml:"exp" json:"exp" valid:"required~ErrorTimestampNotPassed"`
	TransactionId       string                               `yaml:"transactionId" json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
	BankAddress         string                               `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}

func (crossTransferRequest *CrossTransferRequest) SetDefaults() {
	if crossTransferRequest.EncryptedSecretKeys == nil {
		crossTransferRequest.EncryptedSecretKeys = []models.AccountSecretKey{}
	}
}
