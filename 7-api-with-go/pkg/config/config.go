package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	PRODUCTION  = "production"
	DEVELOPMENT = "development"
	STAGING     = "staging"
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

func IsProduction() bool {
	if AppConfig == nil {
		panic("AppConfig Is Nill Please Load Config Before Call This Function")
	}
	return AppConfig.AppEnv == PRODUCTION
}
func IsDevelopment() bool {
	if AppConfig == nil {
		panic("AppConfig Is Nill Please Load Config Before Call This Function")
	}
	return AppConfig.AppEnv == DEVELOPMENT
}

func Load() {
	err := godotenv.Load("envs/dev.env")

	if err != nil {
		fmt.Println("Failed to load ENV")
		panic(err)
	}

	DBName := os.Getenv("DB_NAME")
	DbPassword := os.Getenv("DB_PASS")
	DbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	DbHost := os.Getenv("DB_HOST")
	DBUserName := os.Getenv("DB_USER")
	Port, _ := strconv.Atoi(os.Getenv("PORT"))
	AppEnv := os.Getenv("APP_ENV")
	AppEnv = strings.ToLower(strings.TrimSpace(AppEnv))

	if AppEnv == "" {
		AppEnv = DEVELOPMENT
	}

	AppConfig = &Config{
		DBName:     DBName,
		DbPassword: DbPassword,
		DbPort:     DbPort,
		DbHost:     DbHost,
		DBUserName: DBUserName,
		Port:       Port,
		AppEnv:     AppEnv,
	}
	fmt.Println("Config loaded successfully ✅")
}
