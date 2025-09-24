package main

import (
	"fmt"
	"os"

	"mobin.dev/internals/db"
	"mobin.dev/pkg/config"
)

func main() {
	config.Load()
	mainDb, err := db.Connect(config.AppConfig)

	if err != nil {
		fmt.Printf("❌ Failed to connect to DB: %v\n", err)
		os.Exit(1)
	}

	defer mainDb.Close()
}
