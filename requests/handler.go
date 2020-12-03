package requests

type handler struct {
	fun  string
	name string
}

func (h *handler) Name() string {
	return h.name
}

func (h *handler) Func() string {
	return h.fun
}
