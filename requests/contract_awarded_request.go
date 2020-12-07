package requests

type ContractAwardedRequest struct {
	TransferRequests []TransferRequest `yaml:"transferRequests" json:"transferRequests"`
}
