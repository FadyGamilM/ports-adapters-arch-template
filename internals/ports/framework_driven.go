package ports

type DbPort interface {
	// for each operation to the account we save the current balance to the history table and the operation name
	SaveToHistory(curr_balance float32, operation string) error

	// a method to close the connection
	CloseDbConnection()
}
