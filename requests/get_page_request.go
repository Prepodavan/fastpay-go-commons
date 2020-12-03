package requests

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/remotes/rcc"
)

type GetPageRequest struct {
	currencyCodeOption int
	ccName             string
	methodName         string
	PageSize           int32  `json:"pageSize" valid:"required"`
	Bookmark           string `json:"bookmark"`
}

func (g *GetPageRequest) Handler() rcc.Handler {
	return &handler{
		fun:  g.MethodName(),
		name: g.CcName(),
	}
}

func NewGetPageRequest(pageSize int32, bookmark string, opts ...Option) (r *GetPageRequest) {
	r = &GetPageRequest{PageSize: pageSize, Bookmark: bookmark}
	for _, opt := range opts {
		opt(r)
	}
	return
}

func (g *GetPageRequest) CcName() string {
	return fmt.Sprintf("%s%d", g.ccName, g.CurrencyCodeOption())
}

func (g *GetPageRequest) SetCcName(ccName string) {
	g.ccName = ccName
}

func (g *GetPageRequest) MethodName() string {
	return g.methodName
}

func (g *GetPageRequest) SetMethodName(methodName string) {
	g.methodName = methodName
}

func (g *GetPageRequest) CurrencyCodeOption() int {
	return g.currencyCodeOption
}

func (g *GetPageRequest) SetCurrencyCodeOption(currencyCodeOption int) {
	g.currencyCodeOption = currencyCodeOption
}
