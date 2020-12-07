package requests

type PartialPerformContractRequest struct {
	PerformContractRequest
	ActualAmountInitiator int64 `yaml:"actualAmountInitiator" json:"actualAmountInitiator" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
	ActualAmountAcceptor  int64 `yaml:"actualAmountAcceptor" json:"actualAmountAcceptor" valid:"required,range(0|9223372036854775807)~ErrorAmountNegative"`
}
