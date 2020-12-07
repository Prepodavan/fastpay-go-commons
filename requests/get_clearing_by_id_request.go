package requests

type GetClearingByIdRequest struct {
	Id string `yaml:"id" json:"id" valid:"required"`
}
