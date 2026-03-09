package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "host=localhost port=5432 user=user password=abc123 dbname=habib-ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return nil, err
	}
	return dbCon, nil
}
