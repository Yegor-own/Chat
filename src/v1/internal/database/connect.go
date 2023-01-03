package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	dsn := conf()
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
