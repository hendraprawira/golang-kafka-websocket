package db

import (
	"golang-kafka-sarama-gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:lastchild@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err, "ERR")
	}
	DB = db
	DB.AutoMigrate(&models.Client{})
}
