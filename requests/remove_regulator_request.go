package requests

import "fmt"

type RemoveRegulatorRequest struct {
	AddRegulatorRequest
}

func (request *RemoveRegulatorRequest) GetAddress() string {
	return request.Address
}

func (request *RemoveRegulatorRequest) String() string {
	return fmt.Sprintf("removeRegulator%s", request.Address)
}
