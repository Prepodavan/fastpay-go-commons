package requests

type RemoveArbitratorRequest struct {
	Address string `yaml:"address" json:"address" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
}
