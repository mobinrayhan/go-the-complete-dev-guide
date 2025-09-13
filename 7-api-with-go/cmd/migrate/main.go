package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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

	default:
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
