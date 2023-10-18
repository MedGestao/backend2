package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewConnection() (*sql.DB, error) {
	config := DBConfig{
		Host:     "localhost",
		Port:     5440,
		User:     "postgres",
		Password: "12345",
		Database: "med_gestao",
	}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err

}
