package requests

type GetSafeDealByIdRequest struct {
	Id string `yaml:"id" json:"id" valid:"required"`
}
