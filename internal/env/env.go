package env

import (
	"os"
	"strconv"
)

var (
	HOST      = GetStrEnv("AUTH_HOST", "0.0.0.0")
	PORT      = GetStrEnv("AUTH_PORT", "8080")
	APP_ENV   = GetStrEnv("APP_ENV", "dev")
	LOG_LEVEL = GetStrEnv("LOG_LEVEL", "DEBUG")

	DB_HOST        = GetStrEnv("DB_HOST", "")
	DB_PORT        = GetStrEnv("DB_PORT", "")
	DB_NAME        = GetStrEnv("DB_NAME", "")
	DB_USER        = GetStrEnv("DB_USER", "")
	DB_PASSWORD    = GetStrEnv("DB_PASSWORD", "")
	DB_SSLMODE     = GetStrEnv("DB_SSLMODE", "")
	DB_TIMEZONE    = GetStrEnv("DB_TIMEZONE", "")
	MIGRATIONS_DIR = GetStrEnv("MIGRATIONS_DIR", "")
)

func GetIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); len(value) == 0 {
		return defaultValue
	} else {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		} else {
			return defaultValue
		}
	}
}

func GetStrEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); len(value) == 0 {
		return defaultValue
	} else {
		return value
	}
}
