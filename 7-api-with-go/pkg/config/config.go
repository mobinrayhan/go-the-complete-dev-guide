package config

import (
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
}

var AppConfig *Config

func Load() {

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
	Port, _ := strconv.Atoi(env["Port"])

	AppConfig = &Config{
		DBName:     DBName,
		DbPassword: DbPassword,
		DbPort:     DbPort,
		DbHost:     DbHost,
		DBUserName: DBUserName,
		Port:       Port,
	}
	fmt.Println("Config loaded successfully ✅")
}
