package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     int
	DBPort   int
	User     string
	Password string
	DbName   string
	AppEnv   string
}

const (
	DEVELOPMENT = "development"
	PRODUCTION  = "production"
	STAGE       = "stage"
)

var AppConfig *Config

func IsDevMode() bool {
	if AppConfig == nil {
		panic("Please Call IsDevMode Function After Load Config")
	}
	return AppConfig.AppEnv == DEVELOPMENT
}

func IsProdMode() bool {
	if AppConfig == nil {
		fmt.Println("Please Call IsDevMode Function After Load Config")
		return false
	}
	return AppConfig.AppEnv == PRODUCTION
}

func Load() {
	err := godotenv.Load("envs/.env.development")

	if err != nil {
		log.Fatalf("failed to load env : %v ", err)
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	appEnv := strings.TrimSpace(strings.ToLower(os.Getenv("APP_ENV")))

	if appEnv == "" {
		appEnv = DEVELOPMENT
	}

	AppConfig = &Config{Host: host, Port: port, User: user, DBPort: dbPort, Password: password, DbName: dbName, AppEnv: appEnv}
}
