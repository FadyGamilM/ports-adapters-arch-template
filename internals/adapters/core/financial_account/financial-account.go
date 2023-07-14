package financialaccount

import "log"

type Adapter struct {
}

// Factory Method Pattern
func NewAdapter() *Adapter {
	return &Adapter{}
}

func (adapter *Adapter) Deposite(amount float32) (float32, error) {
	log.Println("DOMAIN LAYER | Deposite() method is called")
	return 1000.0, nil
}

func (adapter *Adapter) Withdraw(amount float32) (float32, error) {
	log.Println("DOMAIN LAYER | Withdraw() method is called")
	return -500.20, nil
}
