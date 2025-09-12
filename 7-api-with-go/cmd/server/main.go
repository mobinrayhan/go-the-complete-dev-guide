package main

import (
	"fmt"

	"mobin.dev/internals/db"
	"mobin.dev/pkg/config"
)

func main() {
	config.Load()
	mainDb, err := db.Connect(config.AppConfig)

	if err != nil {
		fmt.Println("Failed To Connect Db")
		panic(err)
	}
	defer mainDb.Close()
}
