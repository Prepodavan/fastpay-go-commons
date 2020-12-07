package requests

type ContractCompleteRequest struct {
	Id string `yaml:"id" json:"id" valid:"required,uuid"`
}
