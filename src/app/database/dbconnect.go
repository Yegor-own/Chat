package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBParams struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBUrl      string
	DBPort     string
}

func ConnectDB(params DBParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		params.DBHost,
		params.DBUser,
		params.DBPassword,
		params.DBName,
		params.DBPort,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})

}
