package api

import (
	"log"

	"github.com/FadyGamilM/hex/internals/ports"
)

type Adapter struct {
	// the application layer implementation depends on the domain (core) layer interface (port)
	financial_acc ports.FinancialAccountPort

	// the application layer implementation used the framework layer (DI)
	db ports.DbPort
}

func (adapter *Adapter) PerformDeposite(amount float32) (float32, error) {
	log.Println("APPLICATION LAYER | PerformDeposite() method is called")
	// use core layer business logic
	curr_balance, err := adapter.financial_acc.Deposite(amount)
	if err != nil {
		log.Println("APPLICATION LAYER | Error occurred at the core layer while executing the Deposite()")
		return 0, err
	}

	// save to database
	err = adapter.db.SaveToHistory(curr_balance, "DEPOSITE")
	if err != nil {
		log.Printf("APPLICATION LAYER | Error while calling SaveToHIstory() on database : %v \n", err)
		return 0, err
	}

	return curr_balance, nil
}

func (adapter *Adapter) PerformWithdraw(amount float32) (float32, error) {
	log.Println("APPLICATION LAYER | PerformDeposite() method is called")
	curr_balance, err := adapter.financial_acc.Withdraw(amount)
	if err != nil {
		log.Println("APPLICATION LAYER | Error occurred at the core layer while executing the Withdraw()")
		return 0, err
	}
	return curr_balance, nil
}

func NewAdapter(fa ports.FinancialAccountPort, db ports.DbPort) *Adapter {
	return &Adapter{financial_acc: fa, db: db}
}
