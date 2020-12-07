package requests

type GetContractByIdRequest struct {
	Id string `yaml:"id" json:"id" valid:"required"`
}
