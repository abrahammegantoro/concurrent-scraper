package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var AppID string

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    AppID = os.Getenv("APP_ID")
    if AppID == "" {
        log.Fatal("DUMMY_API_APP_ID environment variable is not set")
    }
}
