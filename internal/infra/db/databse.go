package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	host = "localhost"
// 	port = 5432
// 	user = "postgres"
// 	password = "postgres"
// 	dbName = "RunningDB"
// )

func NewDataBaseConnection() *gorm.DB {
	dsn := "host=localhost user=bessa password=bessa dbname=RunningDB port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}
	return db
}
