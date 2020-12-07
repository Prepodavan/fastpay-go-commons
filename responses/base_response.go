package responses

type BaseResponse struct {
	ErrorCode     int    `yaml:"errorCode" json:"errorCode"`
	ErrorMessage  string `yaml:"errorMessage" json:"errorMessage"`
	TransactionId string `yaml:"transactionId" json:"transactionId"`
	SenderAddress string `yaml:"senderAddress" json:"senderAddress"`
}
