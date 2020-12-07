package requests

type CheckSignRequest struct {
	MsgHash string  `yaml:"msgHash" json:"msgHash" valid:"required"`
	Exp     int64   `yaml:"exp" json:"exp" valid:"required~ErrorTimestampNotPassed"`
	Sig     SignDto `yaml:"sig" json:"sig" valid:"required"`
}
