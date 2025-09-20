package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"mobin.dev/pkg/config"
)

func Connect(c *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.DBPort, c.User, c.Password, c.DbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w ", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB : %w", err)
	}

	fmt.Println("DB Successfully connected! ✅")

	return db, nil
}
