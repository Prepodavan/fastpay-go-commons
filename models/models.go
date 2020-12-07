package models

import (
	"sort"

	"github.com/SolarLabRU/fastpay-go-commons/enums/cross-transaction-payload-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/cross-transaction-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-category-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/deal-state-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/deal-transfer-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/invite-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/invoice-state-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/member-deal-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/operation-deal-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/safe-deal-deposit-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/transaction-status-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/transaction-type-enum"
)

type Arbitrator struct {
	Address string `yaml:"address" json:"address"`
	Name    string `yaml:"name" json:"name"`
	DocType string `yaml:"docType" json:"docType"`
}

type Currency struct {
	Code       int                             `yaml:"code" json:"code"`
	Name       string                          `yaml:"name" json:"name"`
	Type       currency_type_enum.CurrencyType `yaml:"type" json:"type"`
	Unit       string                          `yaml:"unit" json:"unit"`
	Symbol     string                          `yaml:"symbol" json:"symbol"`
	Decimals   int                             `yaml:"decimals" json:"decimals"`
	Properties map[string]string               `yaml:"properties" json:"properties"`
	Enabled    bool                            `yaml:"enabled" json:"enabled"`
	DocType    string                          `yaml:"docType" json:"docType"`
}

type TransactionHistory struct {
	TxId          string                                    `yaml:"txId" json:"txId"`
	AddressFrom   string                                    `yaml:"addressFrom" json:"addressFrom"`
	AddressTo     string                                    `yaml:"addressTo" json:"addressTo"`
	TxType        transaction_type_enum.TransactionType     `yaml:"txType" json:"txType"`
	Status        transaction_status_enum.TransactionStatus `yaml:"status" json:"status"`
	Amount        int64                                     `yaml:"amount" json:"amount"`
	CurrencyCode  int                                       `yaml:"currencyCode" json:"currencyCode"`
	ErrorCode     int                                       `yaml:"errorCode" json:"errorCode"`
	ErrorMessage  string                                    `yaml:"errorMessage" json:"errorMessage"`
	Payload       string                                    `yaml:"payload" json:"payload"`
	Timestamp     int64                                     `yaml:"timestamp" json:"timestamp"`
	TransactionId string                                    `yaml:"transactionId" json:"transactionId"`
	SenderAddress string                                    `yaml:"senderAddress" json:"senderAddress"`
	Data          string                                    `yaml:"data" json:"data"`
	InvoiceNumber string                                    `yaml:"invoiceNumber" json:"invoiceNumber"`
	OrdinalNumber int                                       `yaml:"ordinalNumber" json:"ordinalNumber"`
}

type TransactionHistoryEvent struct {
	History TransactionHistory `yaml:"history" json:"history"`
}

type ExecutedTransactionCurrencyExchangeContractItem struct {
	From          string                                                 `yaml:"from" json:"from"`
	To            string                                                 `yaml:"to" json:"to"`
	CurrencyCode  int                                                    `yaml:"currencyCode" json:"currencyCode"`
	Amount        int64                                                  `yaml:"amount" json:"amount"`
	InvoiceNumber string                                                 `yaml:"invoiceNumber" json:"invoiceNumber"`
	Payload       cross_transaction_payload_enum.CrossTransactionPayload `yaml:"payload" json:"payload"`
}

type CrossTransactionHistory struct {
	Routes               []CurrencyContractRoutingItem                        `yaml:"routes" json:"routes"`
	AddressFrom          string                                               `yaml:"addressFrom" json:"addressFrom"`
	Timestamp            int64                                                `yaml:"timestamp" json:"timestamp"`
	TransactionId        string                                               `yaml:"transactionId" json:"transactionId"`
	Amount               int64                                                `yaml:"amount" json:"amount"`
	Payload              string                                               `yaml:"payload" json:"payload"`
	To                   string                                               `yaml:"to" json:"to"`
	EncryptedSecretKeys  []AccountSecretKey                                   `yaml:"encryptedSecretKeys" json:"encryptedSecretKeys"`
	CurrencyCodeFrom     int                                                  `yaml:"currencyCodeFrom" json:"currencyCodeFrom"`
	CurrencyCodeTo       int                                                  `yaml:"currencyCodeTo" json:"currencyCodeTo"`
	CustomerIdentifier   string                                               `yaml:"customerIdentifier" json:"customerIdentifier"`
	CountryCode          string                                               `yaml:"countryCode" json:"countryCode"`
	Details              []ExecutedTransactionCurrencyExchangeContractItem    `yaml:"details" json:"details"`
	SenderAddress        string                                               `yaml:"senderAddress" json:"senderAddress"`
	BankAddress          string                                               `yaml:"bankAddress" json:"bankAddress"`
	Status               cross_transaction_status_enum.CrossTransactionStatus `yaml:"status" json:"status"`
	TxId                 string                                               `yaml:"txId" json:"txId"`
	ErrorCode            int                                                  `yaml:"errorCode" json:"errorCode"`
	ErrorMessage         string                                               `yaml:"errorMessage" json:"errorMessage"`
	Data                 string                                               `yaml:"data" json:"data"`
	InvoiceNumber        string                                               `yaml:"invoiceNumber" json:"invoiceNumber"`
	TransactionHistories []TransactionHistory                                 `yaml:"transactionHistories" json:"transactionHistories"`
}

type CrossTransactionHistoryEvent struct {
	Data CrossTransactionHistory `yaml:"history" json:"history"`
}

type DepositedBalance struct {
	WithdrawResult
	IssueBankAddress string           `yaml:"issueBankAddress" json:"issueBankAddress"`
	BanksBalance     map[string]int64 `yaml:"banksBalance" json:"banksBalance"`
}

type WithdrawResult struct {
	AccountAmount    int64  `yaml:"accountAmount" json:"accountAmount"`
	IssueBankAmount  int64  `yaml:"issueBankAmount" json:"issueBankAmount"`
	BanksTotalAmount int64  `yaml:"banksTotalAmount" json:"banksTotalAmount"`
	TxId             string `yaml:"txId" json:"txId"`
}

type ClaimsItem struct {
	CurrencyCode    int    `yaml:"currencyCode" json:"currencyCode" valid:"required,range(0|999)~ErrorCurrencyCodeRange"`
	BankClaims      string `yaml:"bankClaims" json:"bankClaims" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	BankLiabilities string `yaml:"bankLiabilities" json:"bankLiabilities" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount          int64  `yaml:"amount" json:"amount"`
	Unconfirmed     int64  `yaml:"unconfirmed" json:"unconfirmed"`
}

type ClaimsItemDoc struct {
	ClaimsItem
	DocType string `yaml:"docType" json:"docType"`
	Id      string `yaml:"id" json:"id"`
}

type BankInfo struct {
	Address string `yaml:"address" json:"address"`
	Name    string `yaml:"name" json:"name"`
	Bik     string `yaml:"bik" json:"bik"`
}

type ClearingInfo struct {
	Id          string                   `yaml:"id" json:"id"`
	Banks       map[string]*ClearingBank `yaml:"banks" json:"banks"`
	Owner       string                   `yaml:"owner" json:"owner"`
	Claims      int64                    `yaml:"claims" json:"claims"`
	Liabilities int64                    `yaml:"liabilities" json:"liabilities"`
	History     []ClaimsItem             `yaml:"history" json:"history"`
	Netting     map[string]int64         `yaml:"netting" json:"netting"`
	Procedure   []ClaimsItem             `yaml:"procedure" json:"procedure"`
	Created     int64                    `yaml:"created" json:"created"`
	DocType     string                   `yaml:"docType" json:"docType"`
}

func (ci *ClearingInfo) GetSortBanksKeys() []string {
	keys := make([]string, 0)
	for k, _ := range ci.Banks {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func (ci *ClearingInfo) GetSortNettingKeys() []string {
	keys := make([]string, 0)
	for k, _ := range ci.Netting {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

type ClearingBank struct {
	Claims      int64 `yaml:"claims" json:"claims"`
	Liabilities int64 `yaml:"liabilities" json:"liabilities"`
}

type ClaimsAggregate struct {
	Amount      int64    `yaml:"amount" json:"amount"`
	Unconfirmed int64    `yaml:"unconfirmed" json:"unconfirmed"`
	Ids         []string `yaml:"ids" json:"ids"`
}

type ClientBank struct {
	Address                string                                                              `yaml:"address" json:"address"`
	BankDisplayName        string                                                              `yaml:"bankDisplayName" json:"bankDisplayName"`
	State                  state_enum.State                                                    `yaml:"state" json:"state"`
	CountryCode            string                                                              `yaml:"countryCode" json:"countryCode"`
	Owner                  string                                                              `yaml:"owner" json:"owner"`
	AvailableContractTypes []currency_exchange_contract_type_enum.CurrencyExchangeContractType `yaml:"availableContractTypes" json:"availableContractTypes"`
	Params                 map[string]string                                                   `yaml:"params" json:"params"`
	DocType                string                                                              `yaml:"docType" json:"docType"`
}

func (cb *ClientBank) GetSortParamsKeys() []string {
	keys := make([]string, 0)
	for k, _ := range cb.Params {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

type CurrencyExchangeContractMutable struct {
	Id                   string                                                                    `yaml:"id" json:"id" valid:"required"`
	AddressAccountSell   string                                                                    `yaml:"addressAccountSell" json:"addressAccountSell" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	AddressAccountBuy    string                                                                    `yaml:"addressAccountBuy" json:"addressAccountBuy" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	AddressCommission    string                                                                    `yaml:"addressCommission" json:"addressCommission" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyInfoSell     CurrencyInfo                                                              `yaml:"currencyInfoSell" json:"currencyInfoSell"`
	CurrencyInfoBuy      CurrencyInfo                                                              `yaml:"currencyInfoBuy" json:"currencyInfoBuy"`
	Category             currency_exchange_contract_category_enum.CurrencyExchangeContractCategory `yaml:"category" json:"category"`
	Type                 currency_exchange_contract_type_enum.CurrencyExchangeContractType         `yaml:"type" json:"type"`
	Price                float64                                                                   `yaml:"price" json:"price" valid:"range(0|9223372036854775807)"`
	FractionalCommission float64                                                                   `yaml:"fractionalCommission" json:"fractionalCommission" valid:"range(0|1)"`
	MaxCommission        int64                                                                     `yaml:"maxCommission" json:"maxCommission" valid:"range(0|9223372036854775807)"`
	MinAmount            int64                                                                     `yaml:"minAmount" json:"minAmount" valid:"range(0|9223372036854775807)"`
	MaxAmount            int64                                                                     `yaml:"maxAmount" json:"maxAmount" valid:"range(0|9223372036854775807)"`
	StartDate            int64                                                                     `yaml:"startDate" json:"startDate" valid:"range(0|9223372036854775807)"`
	EndDate              int64                                                                     `yaml:"endDate" json:"endDate" valid:"range(0|9223372036854775807)"`
}

type BaseEvent struct {
	ChaincodeName string `yaml:"chaincodeName" json:"chaincodeName"`
	FunctionName  string `yaml:"functionName" json:"functionName"`
}

type EventBatch struct {
	Events []EventBatchItem `yaml:"events" json:"events"`
}

type EventBatchItem struct {
	EventName string      `yaml:"eventName" json:"eventName"`
	Data      interface{} `yaml:"data" json:"data"`
}

type CurrencyEvent struct {
	BaseEvent
	Data Currency `yaml:"data" json:"data"`
}

type AccountEvent struct {
	BaseEvent
	Data Account `yaml:"data" json:"data"`
}

type ArbitratorEvent struct {
	BaseEvent
	Data Arbitrator `yaml:"data" json:"data"`
}

type BankEvent struct {
	BaseEvent
	Data Bank `yaml:"data" json:"data"`
}

type AvailablePlatformsEvent struct {
	BaseEvent
	Data string `yaml:"data" json:"data"`
}

type TransactionEvent struct {
	BaseEvent
	Data []TransactionHistory `yaml:"data" json:"data"`
}

type CrossTransactionEvent struct {
	BaseEvent
	Data CrossTransactionHistory `yaml:"data" json:"data"`
}

type ClaimsItemResponse struct {
	CurrencyCode    int      `yaml:"currencyCode" json:"currencyCode"`
	BankClaims      BankInfo `yaml:"bankClaims" json:"bankClaims"`
	BankLiabilities BankInfo `yaml:"bankLiabilities" json:"bankLiabilities"`
	Amount          int64    `yaml:"amount" json:"amount"`
	Unconfirmed     int64    `yaml:"unconfirmed" json:"unconfirmed"`
}

type ClearingData struct {
	Id          string               `yaml:"id" json:"id"`
	Owner       string               `yaml:"owner" json:"owner"`
	Claims      int64                `yaml:"claims" json:"claims"`
	Liabilities int64                `yaml:"liabilities" json:"liabilities"`
	History     []ClaimsItemResponse `yaml:"history" json:"history"`
	Netting     []AmountOfBank       `yaml:"netting" json:"netting"`
	Procedure   []ClaimsItemResponse `yaml:"procedure" json:"procedure"`
	Created     int64                `yaml:"created" json:"created"`
}

type ClearingEvent struct {
	BaseEvent
	Data ClearingData `yaml:"data" json:"data"`
}

type ClaimsEvent struct {
	BaseEvent
	Data []ClaimsItem `yaml:"data" json:"data"`
}

type ClientBankEvent struct {
	BaseEvent
	Data ClientBank `yaml:"data" json:"data"`
}

type CustomerEvent struct {
	BaseEvent
	Data Customer `yaml:"data" json:"data"`
}

type ContractEvent struct {
	BaseEvent
	Data CurrencyExchangeContract `yaml:"data" json:"data"`
}

type SetBalanceAccountParam struct {
	Address     string
	AddressBank string
	Value       int64
	Operation   string
	TxId        string
}

type Invoice struct {
	Number       string                          `yaml:"number" json:"number"`
	CurrencyCode int                             `yaml:"currencyCode" json:"currencyCode"`
	Amount       int64                           `yaml:"amount" json:"amount"`
	Description  string                          `yaml:"description" json:"description"`
	Recipient    string                          `yaml:"recipient" json:"recipient"`
	Payer        string                          `yaml:"payer" json:"payer"`
	State        invoice_state_enum.InvoiceState `yaml:"state" json:"state"`
	Created      int64                           `yaml:"created" json:"created"`
	Updated      int64                           `yaml:"updated" json:"updated"`
	ErrorCode    int                             `yaml:"errorCode" json:"errorCode"`
	Owner        string                          `yaml:"owner" json:"owner"`
	DocType      string                          `yaml:"docType" json:"docType"`
}

type InvoiceEvent struct {
	BaseEvent
	Data Invoice `yaml:"data" json:"data"`
}

type LimitsAccount struct {
	Daily   int64 `yaml:"daily" json:"daily"`
	Monthly int64 `yaml:"monthly" json:"monthly"`
}

type Deal struct {
	Id                    string                            `yaml:"id" json:"id"`
	OfferId               string                            `yaml:"offerId" json:"offerId"`
	SenderBank            string                            `yaml:"senderBank" json:"senderBank"`
	Owner                 string                            `yaml:"owner" json:"owner"`
	State                 deal_state_enum.DealState         `yaml:"state" json:"state"`
	Terms                 TermsDeal                         `yaml:"terms" json:"terms"`
	ActualTerms           TermsDeal                         `yaml:"actualTerms" json:"actualTerms"`
	Invitations           map[string]*Invitation            `yaml:"invitations" json:"invitations"`
	Participants          map[string]*Participant           `yaml:"participants" json:"participants"`
	TermsContractConclude map[string]*TermsContractConclude `yaml:"termsContractConclude" json:"termsContractConclude"`
	NecessaryTransfers    map[string]*TransferSafeDeal      `yaml:"necessaryTransfers" json:"necessaryTransfers"`
	DepositHistory        []DepositSafeDealHistory          `yaml:"depositHistory" json:"depositHistory"`
}

type TransferSafeDeal struct {
	AddressFrom  string           `yaml:"addressFrom" json:"addressFrom"`
	AddressTo    string           `yaml:"addressTo" json:"addressTo"`
	CurrencyInfo CurrencyDealInfo `yaml:"currencyInfo" json:"currencyInfo"`
	Amount       int64            `yaml:"amount" json:"amount" valid:"optional,range(0|9223372036854775807)"`
}

type TermsContractConclude struct {
	AddressFrom               string                               `yaml:"addressFrom" json:"addressFrom"`
	AddressTo                 string                               `yaml:"addressTo" json:"addressTo"`
	EscrowAddress             string                               `yaml:"escrowAddress" json:"escrowAddress"`
	MemberTypeTo              member_deal_type_enum.MemberDealType `yaml:"memberTypeTo" json:"memberTypeTo"`
	TxId                      string                               `yaml:"txId" json:"txId"`
	IsComplete                bool                                 `yaml:"isComplete" json:"isComplete"`
	IsCompleteByCrossTransfer bool                                 `yaml:"isCompleteByCrossTransfer" json:"isCompleteByCrossTransfer"`
	IsInvoiceCreate           bool                                 `yaml:"isInvoiceCreate" json:"isInvoiceCreate"`
	CurrencyInfo              CurrencyDealInfo                     `yaml:"currencyInfo" json:"currencyInfo"`
	ObligatoryAmount          int64                                `yaml:"obligatoryAmount" json:"obligatoryAmount"`
	CurrentAmount             int64                                `yaml:"currentAmount" json:"currentAmount"`
	NeedAmount                int64                                `yaml:"needAmount" json:"needAmount"`
}

type Invitation struct {
	AddressFrom        string                               `yaml:"addressFrom" json:"addressFrom"`
	InviteAddress      string                               `yaml:"inviteAddress" json:"inviteAddress"`
	InviteCurrencyInfo CurrencyDealInfo                     `yaml:"inviteCurrencyInfo" json:"inviteCurrencyInfo"`
	Created            int64                                `yaml:"created" json:"created"`
	InviteStatus       invite_status_enum.InviteStatus      `yaml:"inviteStatus" json:"inviteStatus"`
	MemberType         member_deal_type_enum.MemberDealType `yaml:"memberType" json:"memberType"`
}

type Participant struct {
	Address        string                                       `yaml:"address" json:"address"`
	CurrencyInfo   CurrencyDealInfo                             `yaml:"currencyInfo" json:"currencyInfo"`
	MemberType     member_deal_type_enum.MemberDealType         `yaml:"memberType" json:"memberType"`
	Created        int64                                        `yaml:"created" json:"created"`
	WasInvited     bool                                         `yaml:"wasInvited" json:"wasInvited"`
	TransferStatus deal_transfer_status_enum.DealTransferStatus `yaml:"transferStatus" json:"transferStatus"`
}

type TermsDeal struct {
	AddressInitiator       string                                     `yaml:"addressInitiator" json:"addressInitiator" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyInfoInitiator  CurrencyDealInfo                           `yaml:"currencyInfoInitiator" json:"currencyInfoInitiator" valid:"required"`
	AmountInitiator        int64                                      `yaml:"amountInitiator" json:"amountInitiator" valid:"range(0|9223372036854775807)"`      // Сумму которую инициатор отдает
	OperationTypeInitiator operation_deal_type_enum.OperationDealType `yaml:"operationTypeInitiator" json:"operationTypeInitiator" valid:"required,range(0|2)"` // В каком виде инициатор отдает указанную сумму
	Price                  float64                                    `yaml:"price" json:"price" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	MinAmount              int64                                      `yaml:"minAmount" json:"minAmount" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	MaxAmount              int64                                      `yaml:"maxAmount" json:"maxAmount" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	AddressAcceptor        string                                     `yaml:"addressAcceptor" json:"addressAcceptor" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
	CurrencyInfoAcceptor   CurrencyDealInfo                           `yaml:"currencyInfoAcceptor" json:"currencyInfoAcceptor" valid:"required"`
	OperationTypeAcceptor  operation_deal_type_enum.OperationDealType `yaml:"operationTypeAcceptor" json:"operationTypeAcceptor" valid:"required,range(0|2)"`
	AmountAcceptor         int64                                      `yaml:"amountAcceptor" json:"amountAcceptor" valid:"optional,range(0|9223372036854775807)~ErrorAmountNegative"`
}

// TODO: В случае, если CurrencyDealInfo не будет в дальнейшем расходится с CurrencyInfo - поправить сущности безопасных сделок
type CurrencyDealInfo struct {
	Code int    `yaml:"code" json:"code" valid:"range(0|999)~ErrorCurrencyCodeRange"`
	Name string `yaml:"name" json:"name" valid:"required"`
	Unit string `yaml:"unit" json:"unit"`
}

type CurrencyInfo struct {
	Code   int    `yaml:"code" json:"code" valid:"range(0|999)~ErrorCurrencyCodeRange"`
	Symbol string `yaml:"symbol" json:"symbol"`
	Unit   string `yaml:"unit" json:"unit"`
}

type SafeDealEvent struct {
	BaseEvent
	Data DealResponseData `yaml:"data" json:"data"`
}

type SafeDealDeposit struct {
	SafeDealId         string                   `yaml:"safeDealId" json:"safeDealId"`
	Deposits           []DepositDetails         `yaml:"deposits" json:"deposits"`
	CurrentBalance     []SetBalanceAccountParam `yaml:"currentBalance" json:"currentBalance"`
	ForCompleteBalance []SetBalanceAccountParam `yaml:"forCompleteBalance" json:"forCompleteBalance"`
	AddressTo          string                   `yaml:"addressTo" json:"addressTo"`
	CurrencyInfo       CurrencyDealInfo         `yaml:"currencyInfo" json:"currencyInfo"`
	IsComplete         bool                     `yaml:"isComplete" json:"isComplete"`
	CurrentAmount      int64                    `yaml:"currentAmount" json:"currentAmount"`
	NeedAmount         int64                    `yaml:"needAmount" json:"needAmount"`
}

type DepositDetails struct {
	AddressFrom string `yaml:"addressFrom" json:"addressFrom"`
	Amount      int64  `yaml:"amount" json:"amount"`
	TxID        string `yaml:"txID" json:"txID"`
}

type DepositSafeDealHistory struct {
	AddressFrom   string                                          `yaml:"addressFrom" json:"addressFrom"`
	AddressTo     string                                          `yaml:"addressTo" json:"addressTo"`
	CurrencyInfo  CurrencyDealInfo                                `yaml:"currencyInfo" json:"currencyInfo"`
	Amount        int64                                           `yaml:"amount" json:"amount"`
	Type          safe_deal_deposit_type_enum.SafeDealDepositType `yaml:"type" json:"type"`
	TxID          string                                          `yaml:"txID" json:"txID"`
	Timestamp     int64                                           `yaml:"timestamp" json:"timestamp"`
	OrdinalNumber int                                             `yaml:"ordinalNumber" json:"ordinalNumber"`
}

type DealResponseData struct {
	Id                    string                    `yaml:"id" json:"id"`
	OfferId               string                    `yaml:"offerId" json:"offerId"`
	Owner                 string                    `yaml:"owner" json:"owner"`
	State                 deal_state_enum.DealState `yaml:"state" json:"state"`
	Terms                 TermsDeal                 `yaml:"terms" json:"terms"`
	ActualTerms           TermsDeal                 `yaml:"actualTerms" json:"actualTerms"`
	Invitations           []Invitation              `yaml:"invitations" json:"invitations"`
	Participants          []Participant             `yaml:"participants" json:"participants"`
	TermsContractConclude []TermsContractConclude   `yaml:"termsContractConclude" json:"termsContractConclude"`
	NecessaryTransfers    []TransferSafeDeal        `yaml:"necessaryTransfers" json:"necessaryTransfers"`
	DepositHistory        []DepositSafeDealHistory  `yaml:"depositHistory" json:"depositHistory"`
}

type AccountBalance struct {
	AccountBalanceAddress     string `yaml:"accountBalanceAddress" json:"accountBalanceAddress"`
	AccountBalanceAddressBank string `yaml:"accountBalanceAddressBank" json:"accountBalanceAddressBank"`
	Op                        string `yaml:"op" json:"op"`
	Amount                    int64  `yaml:"amount" json:"amount"`
	TxID                      string `yaml:"txID" json:"txID"`
	DocType                   string `yaml:"docType" json:"docType"`
}

type TransferHistoryLimit struct {
	TransferHistoryAddress string `yaml:"transferHistoryLimitAddress" json:"transferHistoryLimitAddress"`
	YearMonth              int    `yaml:"yearMonth" json:"yearMonth"`
	Day                    int    `yaml:"day" json:"day"`
	Value                  int64  `yaml:"value" json:"value"`
	TxID                   string `yaml:"txID" json:"txID"`
	DocType                string `yaml:"docType" json:"docType"`
}

type AccountSecretKey struct {
	AddressTo          string `yaml:"addressTo" json:"addressTo"`
	EncryptedSecretKey string `yaml:"encryptedSecretKey" json:"encryptedSecretKey"`
}
