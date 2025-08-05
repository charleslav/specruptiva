package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SqliteConfig struct {
	LogMode bool
	DbFile  string
}

func InitDb(config SqliteConfig) *gorm.DB {
	db, err := gorm.Open("sqlite3", config.DbFile)
	if err != nil {
		panic(err)
	}
	db.LogMode(config.LogMode)
	return db
}
