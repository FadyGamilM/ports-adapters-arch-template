package db

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

// max number of open connections in the db pool
const maxOpenDbConn = 10

// the max number of connections that can remain open but unused in the conn pool
// if a conn is an idle and this limit is reached, this conn will be closed automatically
const maxIdleDbConn = 5

// the maximum time that a specific conn exists (used or unused) and after this time this conn
// will be destroyed even if its used
const maxDbLifeTime = 5 * time.Minute

const dbTimeOut time.Duration = 3 * time.Second

func NewAdapter(driver_name, data_source_name string) (*Adapter, error) {
	// connect to the database
	db, err := sql.Open(driver_name, data_source_name)
	if err != nil {
		log.Fatalf("FRAMEWORK LAYER | Driven Side | error while connecting to the database : %v \n", err)
	}

	// set pool conn attributes
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	// test the connection if its successffuly done
	err = db.Ping()
	if err != nil {
		log.Fatalf("FRAMEWORK LAYER | Driven Side | I Pinged, Db didnot pong back so connection is no more live : %v \n ", err)
	}

	// if everything is okay return the adapter
	return &Adapter{db: db}, nil
}

func (adapter *Adapter) CloseDbConnection() {
	err := adapter.db.Close()
	if err != nil {
		log.Fatalf("FRAMEWORK LAYER | Driven Side | Error while closing the database connection : %v \n ", err)
	}
}

func (adapter *Adapter) SaveToHistory(curr_balance float32, operation string) error {
	// define a context with time out
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `INSERT INTO accounts (date, curr_balance, operation) VALUES ($1, $2, $3) RETURNING *`

	_, err := adapter.db.ExecContext(ctx, query, time.Now(), curr_balance, operation)
	if err != nil {
		log.Fatalf("FRAMEWORK LAYER | Driven Side | Error while executing the INSERT query to database : %v \n", err)
		return err
	}

	return nil
}
