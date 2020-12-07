package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/models"
)

type CreateApplicationRequest struct {
	Id      string           `yaml:"id" json:"id" valid:"required,uuid"`
	OfferId string           `yaml:"offerId" json:"offerId"`
	Owner   string           `yaml:"owner" json:"owner" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Terms   models.TermsDeal `yaml:"terms" json:"terms" valid:"required"`
	MsgHash string           `yaml:"msgHash" json:"msgHash" valid:"required"`
	Sig     SignDto          `yaml:"sig" json:"sig" valid:"required"`
	Exp     int64            `yaml:"exp" json:"exp" valid:"required~ErrorTimestampNotPassed"`
}
