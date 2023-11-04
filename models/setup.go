package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password= dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	err = db.AutoMigrate(&MyProjects{}, &Blog{}, &Certificate{}, &User{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	DB = db
	log.Println("🚀 Connected Successfully to the Database")
}