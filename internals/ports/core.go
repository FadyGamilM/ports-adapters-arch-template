package ports

// for any adapter need to work with the core port, you have to implement these 2 methods
type FinancialAccountPort interface {
	Deposite(float32) (float32, error)
	Withdraw(float32) (float32, error)
}
