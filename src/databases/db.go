package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.ViperEnv("DB_USER"),
		config.ViperEnv("DB_PASS"),
		config.ViperEnv("DB_HOST"),
		config.ViperEnv("DB_PORT"),
		config.ViperEnv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		panic("Could not connect with the database!")
	}

	log.Println("Connected to database successfully")

	// DEBUG ON CONSOLE
	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate(
		&models.Config{},
		&models.Adnet{},
		&models.Blacklist{},
		&models.Content{},
		&models.Service{},
		&models.Transaction{},
		&models.Subscription{},
		&models.Retry{},
	)

	// TODO: Seed records

	Database = DbInstance{
		Db: db,
	}
}
