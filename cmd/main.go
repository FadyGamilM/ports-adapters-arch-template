package main

import (
	"fmt"

	core_adapter "github.com/FadyGamilM/hex/internals/adapters/core/financial_account"
)

// This file is responsable for app startup and wiring the dependencies via dependency injection
func main() {
	adapter := core_adapter.NewAdapter()

	res, _ := adapter.Withdraw(50)
	fmt.Println(res)
}
