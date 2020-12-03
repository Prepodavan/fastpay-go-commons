package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/remotes/ccs"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/rcc"
)

type GetByIdentifierRequest struct {
	Identifier string `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
}

func (g *GetByIdentifierRequest) Handler() rcc.Handler {
	const fun = "getByIdentifier"
	return &handler{
		fun:  fun,
		name: ccs.Accounts,
	}
}
