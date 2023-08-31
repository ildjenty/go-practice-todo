package db

import (
	"log"
	"os"

	"app/db/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect () *gorm.DB {
	dsn := os.ExpandEnv("host=${DB_HOSTNAME} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable TimeZone=${DB_TZ}")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("database can'nt connect")
	}

	db.AutoMigrate(&model.Todo{})

	return db
}