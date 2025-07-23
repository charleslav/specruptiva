package main

import "github.com/jinzhu/gorm"

// InitDb is database initialization
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&CueSchema{}) {
		db.CreateTable(&CueSchema{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&CueSchema{})
	}

	return db
}
