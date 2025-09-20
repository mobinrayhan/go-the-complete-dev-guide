package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"mobin.dev/internals/db"
	"mobin.dev/pkg/config"
)

func main() {
	config.Load()

	appEnv := config.AppConfig.AppEnv

	fmt.Println(appEnv)

	mainDb, err := db.Connect(config.AppConfig)
	if err != nil {
		fmt.Println("Failed to connect DB ", err)
		os.Exit(1)
	}

	path := fmt.Sprintf("./internals/db/seed/%s", appEnv)

	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Failed to read dir", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		sql, _ := os.ReadFile(filepath.Join(path, file.Name()))
		if _, err := mainDb.Exec(string(sql)); err != nil {
			log.Fatalf("Failed seeding %s: %v", file.Name(), err)
		}
		fmt.Printf("Seeded: %s\n", file.Name())
	}

}
