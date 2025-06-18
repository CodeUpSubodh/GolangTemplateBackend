package config

import (
    "os"
)

var (
    DB_URL = getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable")
)

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
