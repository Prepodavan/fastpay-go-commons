package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/remotes/ccs"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/rcc"
)

type GetCustomersInfoByIdAndCountryCodeRequest struct {
	Identifier  string `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode string `json:"countryCode" valid:"required~ErrorCountryCodeNotPassed,stringlength(3|3)"`
}

func (g *GetCustomersInfoByIdAndCountryCodeRequest) Handler() rcc.Handler {
	const fun = "getCustomerInfoByIdAndCountryCode"
	return &handler{
		fun:  fun,
		name: ccs.CrossCustomers,
	}
}
