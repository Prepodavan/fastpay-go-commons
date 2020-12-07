package requests

import "fmt"

type CreateCustomerRequest struct {
	TechnicalSignRequest
	Identifier          string `yaml:"identifier" json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CustomerDisplayName string `yaml:"customerDisplayName" json:"customerDisplayName"`
	BankAddress         string `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}

func (request *CreateCustomerRequest) String() string {
	return fmt.Sprintf("createCustomer%s%s%s", request.BankAddress, request.Identifier, request.CustomerDisplayName)
}
