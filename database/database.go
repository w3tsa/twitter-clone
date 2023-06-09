package database

import (
	"fmt"
	"log"
	"os"

	"github.com/w3tsa/twitter-clone/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct { 
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", 
	os.Getenv("DB_USER"), 
	os.Getenv("DB_PASSWORD"), 
	os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err!= nil {
        log.Fatal("Failed to connect to the database: \n", err)
		os.Exit(2)
    }
	
	log.Println("connected to the database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	// db.Exec("DROP TABLE posts")
	// db.Exec("DROP TABLE users")
	db.AutoMigrate(&models.User{}, &models.Post{})
	DB = Dbinstance{Db: db}
}