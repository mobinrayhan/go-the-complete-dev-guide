package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"mobin.dev/internals/db"
	"mobin.dev/pkg/config"
)

func main() {
	config.Load()
	mainDb, dbErr := db.Connect(config.AppConfig)

	if dbErr != nil {
		_ = fmt.Errorf("failed to connect db : %w", dbErr)
	}

	up := flag.Bool("up", false, "Apply all up migrations")
	down := flag.Int("down", 0, "Rollback N migrations (e.g., -down 1)")
	force := flag.Int("force", -1, "Set DB schema to specific version (e.g., -force 1)")
	version := flag.Bool("version", false, "Show The Current Migrate Version")
	seed := flag.Bool("seed", false, "Migrate latest Data To The Database!")
	flag.Parse()

	defer mainDb.Close()

	driver, err := mysql.WithInstance(mainDb, &mysql.Config{})

	if err != nil {
		_ = fmt.Errorf("failed to create migrate driver: %w", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance("file://./internals/db/migrations", "mysql", driver)
	if err != nil {
		_ = fmt.Errorf("failed to create migrate instance: %w", err)
		return
	}

	switch {
	case *up:
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration Up Failed : %v", err)
		}
		fmt.Println("✅ All migrations applied")
	case *down > 0:
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration Down Failed : %v", err)
		}
		fmt.Printf("✅ Rolled Back %d migrations(s)\n", *down)
	case *force >= 0:
		if err := m.Force(*force); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Force migration failed: %v", err)
		}
		fmt.Printf("✅ Forced schema version to %d\n", *force)
	case *version:
		v, dirty, err := m.Version()
		if err != nil && err != migrate.ErrNilVersion {
			log.Fatalf("could not get version : %v ", err)
		}
		if err == migrate.ErrNilVersion {
			fmt.Println("📦 No migrations applied yet")
		} else {
			fmt.Printf("📦 Current version: %d (dirty: %v)\n", v, dirty)
		}
	case *seed:
		seedDir := "./internals/db/seeds/dev"
		if config.IsProduction() {
			seedDir = "./internals/db/seeds/prod"
		}
		files, err := filepath.Glob(seedDir + "/*seed.sql")

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fmt.Println("Running Seed ... :", file)
			content, err := os.ReadFile(file)

			if err != nil {
				log.Fatalf("Error reading file : %v", err)
			}

			if string(content) == "" {
				fmt.Println("⚠️ Seed File Is Empty", file)
			} else {
				_, err = mainDb.Exec(string(content))

				if err != nil {
					log.Fatalf("Failed to run seed %s: %v", file, err)
				}
				fmt.Println("✅ Seed applied idempotently", file)
			}
		}

	default:
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
