package main

import (
	"fmt"

	core_adapter "github.com/FadyGamilM/hex/internals/adapters/core/financial_account"
	"github.com/FadyGamilM/hex/internals/ports"
)

// This file is responsable for app startup and wiring the dependencies via dependency injection
func main() {
	// ! Ports section
	var financial_acc_core ports.FinancialAccountPort

	//! plugin the adapter to the port
	financial_acc_core = core_adapter.NewAdapter()

	res, _ := financial_acc_core.Withdraw(50)
	fmt.Println(res)
}
