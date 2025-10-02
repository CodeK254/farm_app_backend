package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	PublicHost string
	DBUser     string
	DBName     string
	DBPassword string
	DBPort     string
	Port     string
	DBAddress  string
	JWTExpiration  int64
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() DBConfig {
	godotenv.Load()
	return DBConfig{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBPort:     getEnv("DB_PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "farm_app_backend"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION", 3600 * 24 * 7),
		JWTSecret: getEnv("JWT_SECRET", "no-secrets-here"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}