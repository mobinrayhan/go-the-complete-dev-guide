package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"mobin.dev/pkg/config"
)

func Connect(c *config.Config) (*sql.DB, error) {

	cfg := mysql.Config{
		User:                 c.DBUserName,
		Passwd:               c.DbPassword,
		DBName:               c.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", c.DbHost, c.DbPort),
		MultiStatements:      true,
	}
	fmt.Println(cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open DB : %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db : %w", err)
	}

	return db, nil
}
