package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "go-be/models"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	log.Println("DB init")

	os.Remove("db.sqlite3")
	db, err := gorm.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic("failed to connect database")
	}

	if err == nil {
		// db.DB().SetMaxIdleConns(conf.MaxIdleConn)
		DB = db
		db.AutoMigrate(&Dept{}, &Proj{}, &User{})
		return db, err
	}

	return nil, err
}
