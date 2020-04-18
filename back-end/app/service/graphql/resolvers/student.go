package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	//"time"
	"github.com/friday182/ttm-admin/app/model"
	//"github.com/friday182/ttm-admin/app/service/graphql/generated"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gvalid"
	"github.com/jinzhu/gorm"
)

// Query
func (r *queryResolver) GetStudents(ctx context.Context, mentorEmail string) ([]*model.Student, error) {
	var existStudents []*model.Student
	err := r.TtmDb.Where("mentor_email=?", mentorEmail).Find(&existStudents).Error
	return existStudents, err
}

func (r *queryResolver) Student(ctx context.Context, gid string, name string) (*model.Student, error) {
	var student model.Student
	err := r.TtmDb.Where("gid=?", gid).Find(&student).Error
	if err != nil {
		err = r.TtmDb.Where("name=?", name).Find(&student).Error
	}
	return &student, err
}

// Mutation
func (r *mutationResolver) AddStudent(ctx context.Context, mentorEmail string, name string, age int) (*model.Student, error) {
	var student model.Student
	params := map[string]interface{}{
		"email": mentorEmail,
		"name":  name,
		"age":   age,
	}
	rules := map[string]string{
		"email": "required|email|length:6,36",
		"name":  "required|length:3,16",
		"Age":   "between:4,16",
	}
	if e := gvalid.CheckMap(params, rules); e != nil {
		return &student, errors.New("Invalid Input")
	}
	var mentor model.Mentor
	err := r.TtmDb.Where("email=?", mentorEmail).First(&mentor).Error
	if err != nil {
		return &student, errors.New("Mentor Not Exist")
	}
	err = r.TtmDb.Where("mentor_email=? AND name=?", mentorEmail, name).First(&student).Error
	if gorm.IsRecordNotFoundError(err) {
		gid := strconv.FormatInt(time.Now().UnixNano(), 36)
		tmpAssignment := []model.KpAssignment{
			{
				Kp:     "MA8",
				QueIdx: 0,
				NumQue: 10,
				Done:   false,
				Pass:   false,
				Desc:   "Basic three numbers plus",
			},
		}
		tmpSubject := []model.Subject{
			{
				Name:         "Math",
				IsEnabled:    true,
				AiEnabled:    false,
				Level:        1,
				Assignment:   tmpAssignment,
				LastDoneTime: time.Now(),
			},
			{
				Name:         "English",
				IsEnabled:    false,
				AiEnabled:    false,
				Level:        1,
				Assignment:   tmpAssignment,
				LastDoneTime: time.Now(),
			},
			/*{
				Name: "VR",
				IsEnabled: false,
				AiEnabled: false,
				Level:     1,
				Assignment: []model.KpAssignment{},
				LastDoneTime: time.Now(),
			},
			{
				Name: "NVR",
				IsEnabled: false,
				AiEnabled: false,
				Level:     1,
				Assignment: []model.KpAssignment{},
				LastDoneTime: time.Now(),
			},*/
		}
		tmpTest := []model.KpUnit{
			{Kp: "MA2", Weight: 2},
			{Kp: "MA5", Weight: 5},
			{Kp: "MA9", Weight: 9},
			{Kp: "MA1", Weight: 1},
		}
		var d []interface{}
		for _, item := range tmpTest {
			d = append(d, item)
		}
		tmpKpPlan := garray.NewSortedArrayFrom(d, func(v1, v2 interface{}) int {
			if v1.(model.KpUnit).Weight < v2.(model.KpUnit).Weight {
				return 1
			}
			if v1.(model.KpUnit).Weight > v2.(model.KpUnit).Weight {
				return -1
			}
			return 0
		})
		/*
			tmpKpPlan.Add(model.KpUnit{Kp: "MA1", Weight: 2})
			tmpKpPlan.Add(model.KpUnit{Kp: "MA15", Weight: 10})
			tmpKpPlan.Add(model.KpUnit{Kp: "MA15", Weight: 8})
		*/
		tmpMq := model.QuizUnit{
			QuizId:  "MQ1",
			Done:    false,
			Pass:    false,
			Desc:    "Assessment Quiz",
			Percent: 0,
		}
		byteMq, _ := json.Marshal(tmpMq)
		byteKpPlan, _ := tmpKpPlan.MarshalJSON()
		byteSubject, _ := json.Marshal(tmpSubject)
		byteEmpty, _ := json.Marshal([]string{""})
		student = model.Student{
			Gid:             gid,
			MentorEmail:     mentorEmail,
			Name:            name,
			Password:        name[:3] + grand.Digits(3),
			Gender:          "",
			Age:             age,
			Country:         "",
			City:            "",
			School:          "",
			Activity:        100,
			Stickers:        0,
			Stars:           0,
			TaskDoneCount:   0,
			TaskPassCount:   0,
			LogUpdated:      true,
			IsMember:        false,
			MembershipStart: time.Now(),
			MembershipStop:  time.Now(),
			LastLoginTime:   time.Now(),
			Quiz:            byteMq,
			StickerLog:      byteEmpty,
			Contacts:        byteEmpty,
			SubjectState:    byteSubject,
			KpPlan:          byteKpPlan,
		}
		err = r.TtmDb.Save(&student).Error
		if err == nil {
			tmpReviseMap := model.ReviseMap{}
			tmpReviseMap.Gid = gid
			r.TtmDb.Create(&tmpReviseMap)
		}
	} else {
		err = errors.New("Student Exist")
	}
	student.Password = "NA"
	return &student, err
}

func (r *mutationResolver) DelStudent(ctx context.Context, gid string) (bool, error) {
	var student model.Student
	var result bool = false
	err := r.TtmDb.Where("gid=?", gid).First(&student).Error
	if gorm.IsRecordNotFoundError(err) {
		result = false
	} else {
		if err = r.TtmDb.Delete(&student).Error; err != nil {
			result = false
		} else {
			result = true
		}
	}

	return result, err
}
