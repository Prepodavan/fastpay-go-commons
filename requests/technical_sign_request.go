package requests

type TechnicalSignRequest struct {
	TechnicalMsgHash string  `json:"technicalMsgHash"`
	TechnicalAddress string  `json:"technicalAddress" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	TechnicalSig     SignDto `json:"technicalSig"`
}
