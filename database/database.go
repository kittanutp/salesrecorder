package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/kittanutp/salesrecorder/config"
)

// create connection with postgres db
func CreateConnection() *sql.DB {

	if config.Err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", config.PsqlInfo)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")
	// return the connection
	return db
}

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.PsqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Item{}, &Sale{}, &SaleItem{})
	return db
}
