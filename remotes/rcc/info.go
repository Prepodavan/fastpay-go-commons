package rcc

type Info struct {
	CC     string
	Method string
}

func (i *Info) Name() string {
	return i.CC
}

func (i *Info) Func() string {
	return i.Method
}
