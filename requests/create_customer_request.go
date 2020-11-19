package requests

import "fmt"

type CreateCustomerRequest struct {
	TechnicalSignRequest
	Identifier          string `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CustomerDisplayName string `json:"customerDisplayName"`
	BankAddress         string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}

func (request *CreateCustomerRequest) MsgHash() string {
	return fmt.Sprintf("createCustomer%s%s%s", request.BankAddress, request.Identifier, request.CustomerDisplayName)
}
