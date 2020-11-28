package middlewares

type BaseContract struct{}

func (bc *BaseContract) GetIgnoredFunctions() []string {
	return []string{
		"Validate",
		"MustHaveMSPID",
		"MustHaveSenderAddress",
		"SenderAccess",
		"SenderAvailable",
	}
}
