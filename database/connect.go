package database

import (
	. "go-be/models"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	log.Println("DB init")

	os.Remove("db.sqlite3")
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	// db, err := gorm.Open("mysql", "root:test@tcp(127.0.0.1:3306)/db?charset=utf8&parseTime=true&multiStatements=true")

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
