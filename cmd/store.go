package main

import (
	"database/sql"
	"github.com/lib/pq"
)

func openDB(name string) (*sql.DB, error) {
	connector, err := pq.NewConnector(name)

	db := sql.OpenDB(connector)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
