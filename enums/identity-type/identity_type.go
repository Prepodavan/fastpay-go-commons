package identity_type

type IdentityType int

const (
	Undefined IdentityType = iota
	None
	Simple
	Identified
)

func (identityType IdentityType) String() string {
	return [...]string{"Undefined", "None", "Simple", "Identified"}[identityType]
}
