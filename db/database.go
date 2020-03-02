package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DataStore interface {
}

type DB struct {
	*sqlx.DB
}

func OpenDB(connection string) (*Database, error) {
	db := Database{}

	var err error
	db.DB, err = sqlx.Connect("postgres", connection)

	if err != nil {
		return nil, fmt.Errorf("Error connecting to Database: %v", err)
	}

	return &db, nil
}
