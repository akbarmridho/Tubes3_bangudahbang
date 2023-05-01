package configs

import (
	"backend/models"
	"encoding/json"
	"io"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
	once       sync.Once
}

func (database *Database) lazyInit() {
	database.once.Do(func() {
		host := os.Getenv("HOST")
		port := os.Getenv("PORT")
		dbname := os.Getenv("DBNAME")
		username := os.Getenv("USERNAME")
		password := os.Getenv("PASSWORD")

		dsn := "host=" + host
		dsn += " user=" + username
		dsn += " password=" + password
		dsn += " dbname=" + dbname
		dsn += " port=" + port

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
		if err != nil {
			panic("Cannot connect database")
		}
		err = db.Migrator().DropTable(&models.Query{})
		if err != nil {
			panic("Cannot drop query table")
		}

		err = db.AutoMigrate(
			&models.Query{},
			&models.History{},
		)

		if err != nil {
			panic("Cannot perform migration")
		}
		// Open the queries.json file
		file, err := os.Open("./assets/queries.json")
		if err != nil {
			panic("Cannot open queries.json")
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		// Read the data from the file
		data, err := io.ReadAll(file)
		if err != nil {
			panic("Failed to read the queries.json")
		}

		// Unmarshal the JSON data into a slice of Query structs
		var queries []models.Query
		if err := json.Unmarshal(data, &queries); err != nil {
			panic("Failed to unmarshal queries")
		}

		if err := db.Create(&queries).Error; err != nil {
			panic("Failed to seed")
		}

		database.connection = db
	})
}

func (database *Database) GetConnection() *gorm.DB {
	database.lazyInit()
	return database.connection
}

var DB = &Database{}
