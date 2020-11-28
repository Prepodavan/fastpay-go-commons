package requests

type ContractAwardedRequest struct {
	TransferRequests []TransferRequest `json:"transferRequests"`
}
