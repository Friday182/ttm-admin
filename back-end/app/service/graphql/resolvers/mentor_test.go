package resolvers_test

import (
	"context"
	"testing"
	//"encoding/json"
	//"time"
	"github.com/friday182/ttm-go/app/service/graphql/resolvers"
	//"github.com/friday182/ttm-go/app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var q resolvers.Resolver
var c context.Context
var localGid string
var mentorEmail string = "test1@gmail.com"
var mentorPw string = "test"
var studentName string = "testStudent"
var studentGid string = ""

func TestMain(m *testing.M) {
	db, err := gorm.Open("sqlite3", "../../../../ttm.sqlite3")
	if err != nil {
		return
	}
	q.Db = db
	_, err = q.Query().Mentor(c, mentorEmail)
	student, err := q.Mutation().AddStudent(c, mentorEmail, studentName, 9)
	if err != nil {
		student, err = q.Query().Student(c, "", studentName)
	}

	studentGid = student.Gid

	m.Run()
}

func TestMentor(t *testing.T) {
	_, err := q.Query().Mentor(c, mentorEmail)
	if err != nil {
		t.Fatal("Get mentor failed", err)
	}
}
