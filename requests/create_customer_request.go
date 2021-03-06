package requests

type CreateCustomerRequest struct {
	Identifier          string `json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CustomerDisplayName string `json:"customerDisplayName"`
	BankId              string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}
