package rcc

type Handler interface {
	Name() string
	Func() string
}
