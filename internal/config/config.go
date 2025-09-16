package config

import (
	"os"

	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

func Load() (*m.Config){
	return &m.Config{
        PORT:       getEnv("PORT"),
        DB_HOST:    getEnv("DB_HOST"),
        DB_PORT:    getEnv("DB_PORT"),
        DB_USER:    getEnv("DB_USER"),
        DB_PASSWORD:getEnv("DB_PASSWORD"),
        DB_NAME:    getEnv("DB_NAME"),
    }
}

func getEnv(key string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }

	return ""
}