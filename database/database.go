package database

import (
	"fmt"
	"go-ecommerce-v2/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	*gorm.DB
}

func OpenDbConnection() *gorm.DB {

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	databaseUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", host, username, password, dbName, port)

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	helper.PanicError(err)

	return db
}
