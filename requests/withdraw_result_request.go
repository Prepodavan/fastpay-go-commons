package requests

type WithdrawResultRequest struct {
	AddressFrom string `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	TxId        string `yaml:"txId" json:"txId" valid:"required~ErrorTxIdSNotPassed"`
}
