package resolvers_test

import (
	"fmt"
	"testing"
	//"encoding/json"
	//"time"
	//"github.com/friday182/ttm-admin/app/service/graphql/resolvers"
	"github.com/friday182/ttm-admin/app/model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestAddStudent(t *testing.T) {
	var student *model.Student
	student, err := q.Query().Student(c, "", "test")
	if err == nil {
		q.Mutation().DelStudent(c, student.Gid)
	}
	student, err = q.Mutation().AddStudent(c, mentorEmail, "test", 9)
	localGid = student.Gid
	fmt.Println(localGid)
	if localGid == "" || err != nil {
		t.Fatal("Get task log failed", err)
	}
}

func TestStudent(t *testing.T) {
	var student *model.Student
	student, err := q.Query().Student(c, localGid, "test")
	if student.Gid != localGid || err != nil {
		t.Fatal("test Student failed: ", err)
	}
}

func TestGetStudents(t *testing.T) {
	students, err := q.Query().GetStudents(c, mentorEmail)
	if err != nil {
		t.Fatal("test GetStudents failed: ", err)
	}
	ok := false
	for _,item := range students {
		if item.Gid == studentGid {
			ok = true
			break
		}
	}
	if ok == false {
		t.Fatal("test GetStudents failed: ", err)
	}
}
