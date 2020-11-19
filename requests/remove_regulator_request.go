package requests

import "fmt"

type RemoveRegulatorRequest struct {
	AddRegulatorRequest
}

func (request *RemoveRegulatorRequest) MsgHash() string {
	return fmt.Sprintf("removeRegulator%s", request.Address)
}
