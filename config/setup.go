package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require sslrootcert=%s", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort, config.SSLCertif)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	err = db.AutoMigrate(
		&models.MyProjects{},
		&models.Blog{},
		&models.Certificate{},
		&models.User{})
		
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	DB = db
	log.Println("🚀 Connected Successfully to the Database")
}