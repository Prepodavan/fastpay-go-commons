package requests

type GetCustomerRequest struct {
	BankAddress string `yaml:"bankAddress" json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
	Identifier  string `yaml:"identifier" json:"identifier" valid:"required~ErrorIdentifierNotPassed,validHex64~ErrorIdentifierNotFolowingRegex"`
	CountryCode string `yaml:"countryCode" json:"countryCode" valid:"required~ErrorCountryCodeNotPassed,stringlength(3|3)"`
}
