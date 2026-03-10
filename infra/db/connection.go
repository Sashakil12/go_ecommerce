package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(config config.DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Name, config.SSLMode)
}

func NewConnection(config config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(config)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return nil, err
	}
	return dbCon, nil
}
