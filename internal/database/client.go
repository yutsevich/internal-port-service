package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"internal-port-service/internal/entities"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Port{})
	log.Println("Database Migration Completed...")
}
