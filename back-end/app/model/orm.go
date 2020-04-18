package model

import (
	"log"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"github.com/friday182/ttm-admin/app/service/setting"
)

var (
	ttmDb  *gorm.DB
	queDb  *gorm.DB
	quizDb *gorm.DB
	err    error
)

// ConnectDb to init database
func ConnectDb() error {
	d, err := gorm.Open("sqlite3", "../../ttm-go/ttm.sqlite3")
	if err != nil {
		d, err = gorm.Open("sqlite3", "../ttmnow/ttm.sqlite3")
	}
	// ttmDb, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=ttm password=123456 sslmode=disable")
	if err != nil {
		log.Println("[ORM] Error: ", err)
		panic("[ORM] Connect Database Failed ")
	} else {
		ttmDb = d
	}

	q, err := gorm.Open("sqlite3", "./quiz.sqlite3")
	if err != nil {
		q, err = gorm.Open("sqlite3", "../ttmnow/quiz.sqlite3")
	}
	// ttmDb, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=ttm password=123456 sslmode=disable")
	if err != nil {
		log.Println("[ORM] Error: ", err)
		panic("[ORM] Connect Database Failed ")
	} else {
		quizDb = q
	}

	//mode := setting.database.mode
	mode := "debug"

	if mode != "producton" {
		ttmDb.LogMode(true)
		quizDb.LogMode(true)
	}

	log.Println("[ORM] Database is ready.")

	return nil
}

// GetTtmDb to export ttm main database
func GetTtmDb() *gorm.DB {
	return ttmDb
}

// GetQuizdb to export quiz database
func GetQuizdb() *gorm.DB {
	return quizDb
}
