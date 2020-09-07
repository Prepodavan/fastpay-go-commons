package safe_deal_deposit_type_enum

type SafeDealDepositType int

const (
	Undefined SafeDealDepositType = iota
	Topup
	Complete
	Return
)

func (safeDealDepositType SafeDealDepositType) String() string {
	return [...]string{"Undefined", "Topup", "Complete", "Return"}[safeDealDepositType]
}
