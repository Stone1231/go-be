package database

import (
	. "go-be/models"
	"log"
	"os"

	_ "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	log.Println("DB init")
	sqliteFile := "db.sqlite3"
	os.Remove(sqliteFile)
	db, err := gorm.Open(sqlite.Open(sqliteFile), &gorm.Config{})
	// db, err := gorm.Open("mysql", "root:test@tcp(127.0.0.1:3306)/db?charset=utf8&parseTime=true&multiStatements=true")
	// db, err = gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       connectionString,
	// 	DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
	// 	DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
	// 	DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
	// 	DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
	// 	SkipInitializeWithVersion: false, // smart configure based on used version
	// }), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err == nil {
		// db.DB().SetMaxIdleConns(conf.MaxIdleConn)
		DB = db
		db.AutoMigrate(&Dept{}, &Proj{}, &User{})
		db.Debug()
		return db, err
	}

	return nil, err
}
