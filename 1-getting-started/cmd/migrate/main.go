package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"mobin.dev/internals/db"
	"mobin.dev/pkg/config"
)

func main() {
	config.Load()

	up := flag.Bool("up", false, "Run this command for apply all latest migration")
	down := flag.Int("down", 0, "Run this command for apply Down migration")
	reset := flag.Int("reset", -1, "Run this command for apply Reset Dirty State")
	version := flag.Bool("version", false, "Run this command for Getting Version of DB")
	flag.Parse()

	fmt.Println(*up)
	mainDb, err := db.Connect(config.AppConfig)
	if err != nil {
		fmt.Println("Failed to connect DB ", err)
		os.Exit(1)
	}

	driver, err := postgres.WithInstance(mainDb, &postgres.Config{})

	if err != nil {
		fmt.Println("Failed To Create Postgres driver ❌", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./internals/db/migrations", "postgres", driver)

	if err != nil {
		fmt.Println("Failed To Create Postgres instance ❌", err)
	}

	switch {
	case *up:
		err := m.Up()
		if err != nil {
			fmt.Println("Failed to latest migration ❌", err)
		} else {
			fmt.Println("Latest migration successful!")
		}
	case *down > 0:
		err := m.Steps(-*down)
		if err != nil {
			fmt.Println("Failed to down migration ❌", err)
		} else {
			fmt.Printf("Migration Down (%d) Successfully! \n", -*down)
		}
	case *reset >= 0:
		err := m.Force(*reset)
		if err != nil {
			fmt.Println("Failed to reset migration ❌", err)
		} else {
			fmt.Println("Migration Reset Successfully !", *reset)
		}
	case *version:
		version, dirty, err := m.Version()
		if err != nil {
			fmt.Println("Failed to reset migration ❌", err)
		} else {
			fmt.Println("Current Version Of DB : ", version)
			fmt.Println("Dirty Status (Default false)! : ", dirty)
		}
	}
}
