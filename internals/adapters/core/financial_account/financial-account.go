package financialaccount

type Adapter struct {
}

// Factory Method Pattern
func NewAdapter() *Adapter {
	return &Adapter{}
}

func (adapter *Adapter) Deposite(amount float32) (float32, error) {
	return 1000.0, nil
}

func (adapter *Adapter) Withdraw(amount float32) (float32, error) {
	return -500.20, nil
}
