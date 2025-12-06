package core

import (
    "os"
)

type Config struct {
    DatabaseURL string
    JWTSecret   string
    SMTPHost    string
    SMTPPort    string
    SMTPUser    string
    SMTPPass    string
    Port        string
}

func Load() *Config {
    return &Config{
        DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/portofolio?sslmode=disable"),
        JWTSecret:   getEnv("JWT_SECRET", "dev-secret"),
        SMTPHost:    os.Getenv("SMTP_HOST"),
        SMTPPort:    os.Getenv("SMTP_PORT"),
        SMTPUser:    os.Getenv("SMTP_USER"),
        SMTPPass:    os.Getenv("SMTP_PASS"),
        Port:        getEnv("PORT", "8080"),
    }
}

func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}
