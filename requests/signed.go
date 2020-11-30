package requests

type Signed interface {
	TechnicalSign
	BaseMsgHashRequest
}
