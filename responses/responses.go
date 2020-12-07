package responses

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contract-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type BankResponse struct {
	Data models.Bank `yaml:"data" json:"data"`
	BaseResponse
}

type MakeSafeDealDepositResponse struct {
	Data []models.EventBatchItem `yaml:"data" json:"data"`
	BaseResponse
}

type InvoiceCreateResponse struct {
	IsInvoiceCreate bool `yaml:"isInvoiceCreate" json:"isInvoiceCreate"`
	BaseResponse
}

type AccountListResponse struct {
	Data []models.Account `yaml:"data" json:"data"`
	BaseResponse
}

type ArbitratorListResponse struct {
	Data []models.Arbitrator `yaml:"data" json:"data"`
	BaseResponse
}

type Metadata struct {
	FetchedRecordsCount int32  `yaml:"fetchedRecordsCount" json:"fetchedRecordsCount"`
	Bookmark            string `yaml:"bookmark" json:"bookmark"`
}

type AccountPageData struct {
	Metadata Metadata         `yaml:"metadata" json:"metadata"`
	Items    []models.Account `yaml:"items" json:"items"`
}

type AccountPageResponse struct {
	Data AccountPageData `yaml:"data" json:"data"`
	BaseResponse
}

type SenderAddressData struct {
	Address string `yaml:"address" json:"address"`
}

type SenderAddressResponse struct {
	Data SenderAddressData `yaml:"data" json:"data"`
	BaseResponse
}

type BankListResponse struct {
	Data []models.Bank `yaml:"data" json:"data"`
	BaseResponse
}

type BankTotalData struct {
	Total int `yaml:"total" json:"total"`
}

type BankTotalResponse struct {
	Data BankTotalData `yaml:"data" json:"data"`
	BaseResponse
}

type BankPageData struct {
	Metadata Metadata      `yaml:"metadata" json:"metadata"`
	Items    []models.Bank `yaml:"items" json:"items"`
}

type BankPageResponse struct {
	Data BankPageData `yaml:"data" json:"data"`
	BaseResponse
}

type SenderIsBankResponse struct {
	Data bool `yaml:"data" json:"data"`
	BaseResponse
}

type GetAvailablePlatformsResponse struct {
	Data string `yaml:"data" json:"data"`
	BaseResponse
}

type CurrencyListResponse struct {
	Data []models.Currency `yaml:"data" json:"data"`
	BaseResponse
}

type CurrencyPageData struct {
	Metadata Metadata          `yaml:"metadata" json:"metadata"`
	Items    []models.Currency `yaml:"items" json:"items"`
}

type CurrencyPageResponse struct {
	Data CurrencyPageData `yaml:"data" json:"data"`
	BaseResponse
}

type CurrencyResponse struct {
	Data models.Currency `yaml:"data" json:"data"`
	BaseResponse
}

type BankBalanceData struct {
	Liabilities []models.AmountOfBank `yaml:"liabilities" json:"liabilities"`
	Claims      []models.AmountOfBank `yaml:"claims" json:"claims"`
	Issue       int64                 `yaml:"issue" json:"issue"`
	IssueLimit  int64                 `yaml:"issueLimit" json:"issueLimit"` // TODO убрать если не быдет использоваться
}

type BankBalanceResponse struct {
	Data BankBalanceData `yaml:"data" json:"data"`
	BaseResponse
}

type ClearingResultResponse struct {
	Data models.ClearingData `yaml:"data" json:"data"`
	BaseResponse
}

type ClearingListResponse struct {
	Data []models.ClearingData `yaml:"data" json:"data"`
	BaseResponse
}

type ClaimsListResponse struct {
	Data []models.ClaimsItem `yaml:"data" json:"data"`
	BaseResponse
}

type ClearingPageData struct {
	Metadata Metadata              `yaml:"metadata" json:"metadata"`
	Items    []models.ClearingData `yaml:"items" json:"items"`
}

type ClearingPageResponse struct {
	Data ClearingPageData `yaml:"data" json:"data"`
	BaseResponse
}

type ClaimsPageData struct {
	Metadata Metadata            `yaml:"metadata" json:"metadata"`
	Items    []models.ClaimsItem `yaml:"items" json:"items"`
}

type ClaimsPageResponse struct {
	Data ClaimsPageData `yaml:"data" json:"data"`
	BaseResponse
}

type BankClaimsLiabilities struct {
	Claims      map[string]int64 `yaml:"claims" json:"claims"`
	Liabilities map[string]int64 `yaml:"liabilities" json:"liabilities"`
}

type BankClaimsLiabilitiesResponse struct {
	Data BankClaimsLiabilities `yaml:"data" json:"data"`
	BaseResponse
}

type ClientBankItemResponse struct {
	Address                string                                                              `yaml:"address" json:"address"`
	BankDisplayName        string                                                              `yaml:"bankDisplayName" json:"bankDisplayName"`
	AvailableContractTypes []currency_exchange_contract_type_enum.CurrencyExchangeContractType `yaml:"availableContractTypes" json:"availableContractTypes"`
	State                  state_enum.State                                                    `yaml:"state" json:"state"`
	CountryCode            string                                                              `yaml:"countryCode" json:"countryCode"`
	Params                 map[string]string                                                   `yaml:"params" json:"params"`
	Owner                  string                                                              `yaml:"owner" json:"owner"`
}

type ClientBankResponse struct {
	Data ClientBankItemResponse `yaml:"data" json:"data"`
	BaseResponse
}

type ClientBankPageData struct {
	Metadata Metadata                 `yaml:"metadata" json:"metadata"`
	Items    []ClientBankItemResponse `yaml:"items" json:"items"`
}

type ClientBankPageResponse struct {
	Data ClientBankPageData `yaml:"data" json:"data"`
	BaseResponse
}

type ClientBankParam struct {
	Key   string `yaml:"key" json:"key"`
	Value string `yaml:"value" json:"value"`
}

type CustomerPageData struct {
	Metadata Metadata          `yaml:"metadata" json:"metadata"`
	Items    []models.Customer `yaml:"items" json:"items"`
}

type CustomerPageResponse struct {
	Data CustomerPageData `yaml:"data" json:"data"`
	BaseResponse
}

type CustomersListResponse struct {
	Data []models.Customer `yaml:"data" json:"data"`
	BaseResponse
}

type CustomerResponse struct {
	Data models.Customer `yaml:"data" json:"data"`
	BaseResponse
}

type ContractResponse struct {
	Data models.CurrencyExchangeContract `yaml:"data" json:"data"`
	BaseResponse
}

type CurrencyExchangeContractPageData struct {
	Metadata Metadata                          `yaml:"metadata" json:"metadata"`
	Items    []models.CurrencyExchangeContract `yaml:"items" json:"items"`
}

type CurrencyExchangeContractPageResponse struct {
	Data CurrencyExchangeContractPageData `yaml:"data" json:"data"`
	BaseResponse
}

type GetBestRoutesResponse struct {
	Data [][]models.CurrencyContractRoutingItem `yaml:"data" json:"data"`
	BaseResponse
}

type WithdrawResultResponse struct {
	Data *models.WithdrawResult `yaml:"data" json:"data"`
	BaseResponse
}

type WithdrawConfirmResponse struct {
	History models.TransactionHistory `yaml:"history" json:"history"`
	Data    string                    `yaml:"data" json:"data"`
	BaseResponse
}

type WithdrawRejectResponse WithdrawConfirmResponse

type TransferResponse struct {
	History models.TransactionHistory `yaml:"history" json:"history"`
	Data    string                    `yaml:"data" json:"data"`
	BaseResponse
}

type TransferBatchResponse struct {
	Histories []models.TransactionHistory `yaml:"histories" json:"histories"`
	Data      string                      `yaml:"data" json:"data"`
	BaseResponse
}

type TopupResponse struct {
	Data string `yaml:"data" json:"data"`
	BaseResponse
}

type ExecutedTransactionCurrencyExchangeContractData struct {
	Transactions         []models.ExecutedTransactionCurrencyExchangeContractItem `yaml:"transactions" json:"transactions"`
	Commission           int64                                                    `yaml:"commission" json:"commission"`
	AmountInput          int64                                                    `yaml:"amountInput" json:"amountInput"`
	AmountOutput         int64                                                    `yaml:"amountOutput" json:"amountOutput"`
	TransactionHistories []models.TransactionHistory                              `yaml:"transactionHistories" json:"transactionHistories"`
}

type ExecutedTransactionCurrencyExchangeContractResponse struct {
	Data ExecutedTransactionCurrencyExchangeContractData `yaml:"data" json:"data"`
	BaseResponse
}

type CrossTransactionHistoryPageData struct {
	Metadata Metadata                         `yaml:"metadata" json:"metadata"`
	Items    []models.CrossTransactionHistory `yaml:"items" json:"items"`
}

type CrossTransactionHistoryPageResponse struct {
	Data CrossTransactionHistoryPageData `yaml:"data" json:"data"`
	BaseResponse
}

type AccountAddressResponse struct {
	Data string `yaml:"data" json:"data"`
	BaseResponse
}

type GetBankBalanceTotalResponseData struct {
	Bank        models.BankInfo       `yaml:"bank" json:"bank"`
	Liabilities []models.AmountOfBank `yaml:"liabilities" json:"liabilities"`
	Claims      []models.AmountOfBank `yaml:"claims" json:"claims"`
	Issue       int64                 `yaml:"issue" json:"issue"`
	IssueLimit  int64                 `yaml:"issueLimit" json:"issueLimit"`
}

type GetBankBalanceTotalResponse struct {
	Data []GetBankBalanceTotalResponseData `yaml:"data" json:"data"`
	BaseResponse
}

type InvoiceResponse struct {
	Data models.Invoice `yaml:"data" json:"data"`
	BaseResponse
}

type InvoiceListResponse struct {
	Data []models.Invoice `yaml:"data" json:"data"`
	BaseResponse
}

type InvoicePageData struct {
	Metadata Metadata         `yaml:"metadata" json:"metadata"`
	Items    []models.Invoice `yaml:"items" json:"items"`
}

type InvoicePageResponse struct {
	Data InvoicePageData `yaml:"data" json:"data"`
	BaseResponse
}

type LimitData struct {
	Value int64 `yaml:"value" json:"value"`
}

type LimitResponse struct {
	Data LimitData `yaml:"data" json:"data"`
	BaseResponse
}

type AccountLimitsData struct {
	OperationLimit int64 `yaml:"operationLimit" json:"operationLimit"`
	BalanceLimit   int64 `yaml:"balanceLimit" json:"balanceLimit"`
	DailyLimit     int64 `yaml:"dailyLimit" json:"dailyLimit"`
	MonthlyLimit   int64 `yaml:"monthlyLimit" json:"monthlyLimit"`
}

type AccountLimitsResponse struct {
	Data AccountLimitsData `yaml:"data" json:"data"`
	BaseResponse
}

type SafeDealResponse struct {
	Data models.DealResponseData `yaml:"data" json:"data"`
	BaseResponse
}

type SafeDealDepositResponse struct {
	Data models.SafeDealDeposit `yaml:"data" json:"data"`
	BaseResponse
}
