package requests

type CreateArbitratorRequest struct {
	Address string `yaml:"address" json:"address" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
	Name    string `yaml:"name" json:"name" valid:"required"`
}
