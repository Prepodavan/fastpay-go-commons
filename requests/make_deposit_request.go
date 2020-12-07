package requests

type MakeDepositRequest struct {
	Id          string  `yaml:"id" json:"id" valid:"required,uuid"`
	AddressFrom string  `yaml:"addressFrom" json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AddressTo   string  `yaml:"addressTo" json:"addressTo" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount      int64   `yaml:"amount" json:"amount" valid:"required~ErrorAmountNotPassed,range(0|9223372036854775807)~ErrorAmountNegative"`
	MsgHash     string  `yaml:"msgHash" json:"msgHash" valid:"required"`
	Sig         SignDto `yaml:"sig" json:"sig" valid:"required"`
	Exp         int64   `yaml:"exp" json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
