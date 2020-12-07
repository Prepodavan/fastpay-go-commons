package requests

type SignDto struct {
	R string `yaml:"r" json:"r" valid:"validHex62or64~ErrorSigRNotFolowingRegex"`
	S string `yaml:"s" json:"s" valid:"validHex62or64~ErrorSigSNotFolowingRegex"`
	V int    `yaml:"v" json:"v" valid:"range(27|28)~ErrorSigVIncorrect"`
}
