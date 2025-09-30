package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct{
	PublicHost string
	DBUser string
	DBName string
	DBPassword string
	DBPort string
	DBAddress string
}

var Envs = initConfig()

func initConfig() DBConfig {
	godotenv.Load()
	return DBConfig{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		DBPort: getEnv("DB_PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "farm_app_backend"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value;
	}

	return fallback;
}