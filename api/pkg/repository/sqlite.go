package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

const (
	TableUsers     = "Users"
	TableProducts  = "Products"
	TableLocations = "Locations"
)

var (
	MutexWalletRead   sync.Mutex
	MutexWalletWrite  sync.Mutex
	MutexWalletDelete sync.Mutex
)

type Storage struct {
	db *sql.DB
}

func NewStorageSQL(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func NewSqliteDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./db/telegram.db")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", err)
	}

	return db, nil
}
