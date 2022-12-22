package database

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() (*DBParams, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}

	params := DBParams{
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBHost:     os.Getenv("DATABASE_HOST"),
		DBUrl:      os.Getenv("DATABASE_URL"),
		DBPort:     os.Getenv("DATABASE_PORT"),
	}

	return &params, nil
}
