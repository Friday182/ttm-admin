package model

import (
	"log"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"github.com/friday182/ttm-go/app/service/setting"
)

var (
	db *gorm.DB
	qdb *gorm.DB
	err error
)

func ConnectDb() error {
	d, err := gorm.Open("sqlite3", "./ttm.sqlite3")
	// db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=ttm password=123456 sslmode=disable")
	if err != nil {
		log.Println("[ORM] Error: ", err)
		panic("[ORM] Connect Database Failed ")
	}else{
		db = d
	}

	q, err := gorm.Open("sqlite3", "./qdb.sqlite3")
	// db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=ttm password=123456 sslmode=disable")
	if err != nil {
		log.Println("[ORM] Error: ", err)
		panic("[ORM] Connect Database Failed ")
	}else{
		qdb = q
	}

	//mode := setting.database.mode
	mode := "debug"

	if mode != "producton" {
		db.LogMode(true)
		qdb.LogMode(true)
	}

	log.Println("[ORM] Database is ready.")

	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetQdb() *gorm.DB {
	return qdb
}
