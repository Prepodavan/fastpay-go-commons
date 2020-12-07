package requests

type handler struct {
	fun  string
	name string
}

func (h *handler) Name() string {
	return h.GetName()
}

func (h *handler) Func() string {
	return h.GetFun()
}
