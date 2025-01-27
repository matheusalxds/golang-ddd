package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	User     = "root"
	Password = "root"
	Name     = "fx-fiber-db"
	Port     = "5432"
)

func ConnectDatabase() *gorm.DB {
	connectinString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password%s sslmode=disable", Host, Port, User, Name, Password)
	db, err := gorm.Open(postgres.Open(connectinString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
