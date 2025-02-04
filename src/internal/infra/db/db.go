package db

import (
	"fmt"
	"go-fx-project/src/internal/infra/env"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildDNS(env env.EnvLoader) string {
	envs := env.LoadEnv()

	host := envs.DBHost
	port := envs.DBPort
	user := envs.DBUser
	password := envs.DBPasss
	dbname := envs.DBName

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func ConnectDatabase(env env.EnvLoader) *gorm.DB {
	db, err := gorm.Open(postgres.Open(BuildDNS(env)), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
