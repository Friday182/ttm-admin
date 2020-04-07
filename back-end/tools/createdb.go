package main

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/friday182/ttm-admin/app/model"
)

func main() {
	//createTtmDb()
	createQuestionDb()
}

func createTtmDb () {
	ttm, err := gorm.Open("sqlite3", "../ttm.sqlite3")
	if err != nil {
		return
	}
	defer ttm.Close()

	tables := []interface{}{
		&model.User{},
		&model.Mentor{},
		&model.Student{},
		&model.Question{},
		&model.KpDescription{},
		&model.TaskLog{},
		&model.KpState{},
		&model.Quiz{},
		&model.KpMap{},
		&model.QuizLog{},
		&model.QuizReport{},
		&model.ReviseMap{},
	}
	
	ttm.DropTableIfExists(tables...)
	ttm.CreateTable(tables...)

	// Add default test user
	u := model.Mentor{}
	u.Email = "test@gmail.com"
	u.Password = "test"
	ttm.Save(&u)

	// Add superuser admin
	admin := model.User{}
	admin.Username = "Phoenix"
	admin.Password = "jiangbo007"
	ttm.Save(&admin)
}

func createQuestionDb () {
	db, err := gorm.Open("sqlite3", "../qdb.sqlite3")
	if err != nil {
		return
	}
	defer db.Close()

	tables := []interface{}{
		&model.Question{},
		&model.KpDescription{},
		&model.Quiz{},
		&model.KpMap{},
	}

	db.DropTableIfExists(tables...)
	db.CreateTable(tables...)
}
