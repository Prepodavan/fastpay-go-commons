package requests

type ContractCompleteRequest struct {
	Id string `json:"id" valid:"required,uuid"`
}
