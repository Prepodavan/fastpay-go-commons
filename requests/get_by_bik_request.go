package requests

type GetByBikRequest struct {
	Bik string `yaml:"bik" json:"bik" valid:"required"`
}
