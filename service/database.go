package service

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataStore interface {
	CreateAccount(acc Account) error
	GetAccount(email string) (Account, error)
}

type DB struct {
	*sqlx.DB
}

func OpenDB(connection string) (*DB, error) {

	db, err := sqlx.Connect("postgres", connection)

	if err != nil {
		return nil, fmt.Errorf("Error connecting to Database: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
