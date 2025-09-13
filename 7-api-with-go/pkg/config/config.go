package config

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName     string
	DbPassword string
	DbPort     int
	DbHost     string
	DBUserName string
	Port       int
	AppEnv     string
}

var AppConfig *Config

func Load() {
	appMode := flag.String("mode", "development", "Mode Of Your App, It Should Be Dev Or Production!")
	flag.Parse()
	env, err := godotenv.Read()

	if err != nil {
		fmt.Println("Failed to load ENV")
		panic(err)
	}

	DBName := env["DB_NAME"]
	DbPassword := env["DB_PASS"]
	DbPort, _ := strconv.Atoi(env["DB_PORT"])
	DbHost := env["DB_HOST"]
	DBUserName := env["DB_USER"]
	Port, _ := strconv.Atoi(env["PORT"])

	AppConfig = &Config{
		DBName:     DBName,
		DbPassword: DbPassword,
		DbPort:     DbPort,
		DbHost:     DbHost,
		DBUserName: DBUserName,
		Port:       Port,
		AppEnv:     *appMode,
	}
	fmt.Println("Config loaded successfully ✅", *appMode)
}
