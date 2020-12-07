package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/remotes/ccs"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/rcc"
)

type GetCustomersInfoByIdAndCountryCodeRequest struct {
	Identifier  string `yaml:"identifier" json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode string `yaml:"countryCode" json:"countryCode" valid:"required~ErrorCountryCodeNotPassed,stringlength(3|3)"`
}

func (g *GetCustomersInfoByIdAndCountryCodeRequest) Handler() rcc.Handler {
	const fun = "getCustomerInfoByIdAndCountryCode"
	return &handler{
		fun:  fun,
		name: ccs.CrossCustomers,
	}
}
