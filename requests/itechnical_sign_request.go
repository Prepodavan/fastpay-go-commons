package requests

type TechnicalSign interface {
	GetTechnicalMsgHash() string
	SetTechnicalMsgHash(TechnicalMsgHash string)
	GetTechnicalAddress() string
	SetTechnicalAddress(TechnicalAddress string)
	GetTechnicalSig() SignDto
	SetTechnicalSig(TechnicalSig SignDto)
}
