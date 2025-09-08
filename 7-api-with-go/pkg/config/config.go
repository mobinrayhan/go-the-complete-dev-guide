package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	Port := os.Getenv("PORT")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	return Config{
		Port:       Port,
		DbHost:     DbHost,
		DbPort:     DbPort,
		DbUser:     DbUser,
		DbPassword: DbPassword,
		DbName:     DbName,
	}

}
