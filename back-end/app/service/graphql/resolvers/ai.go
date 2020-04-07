package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"encoding/json"
	"time"

	"github.com/friday182/ttm-admin/app/model"
	"github.com/friday182/ttm-admin/app/service/graphql/generated"
	"github.com/jinzhu/gorm"
	//"github.com/gogf/gf/util/gconv"
)

var subjectMap map[string]int = map[string]int{
	"M": 0,
	"E": 1,
}

type QueUnit struct {
	QueIdx    int      `json:"queIdx"`
	IsCorrect bool     `json:"isCorrect"`
	StdSec    int      `json:"stdSec"`
	UsedSec   int      `json:"usedSec"`
	Tags      []string `json:"tags"`
}

type DetailsLog struct {
	Kp         string    `json:"kp" gorm:"not null"`
	NumQue     int       `json:"numQue"`
	CorrectQue int       `json:"correctQue"`
	MaxSec     int       `json:"maxSec"`
	MinSec     int       `json:"minSec"`
	ResultList []QueUnit `json:"resultList"`
	TotalSec   int       `json:"totalSec"`
	StdSec     int       `json:"stdSec"`
	WrongList  []int
	CorrectPer int
}

// Query
/*
func (r *queryResolver) CheckNewLog(ctx context.Context, gid string) (bool, error) {
	var student model.Student
	err := r.Db.Where("gid = ?", gid).First(&student).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, err
	}
	return student.LogUpdated, nil
}
*/
func (r *queryResolver) GetTaskLogs(ctx context.Context, gid string, startDate string, startID int, numLog int) (*generated.TaskLogOutput, error) {
	var logs []*model.TaskLog
	output := generated.TaskLogOutput{}
	var student model.Student
	err := r.Db.Where("gid = ?", gid).First(&student).Error
	if gorm.IsRecordNotFoundError(err) {
		return &output, err
	}
	if student.LogUpdated == true {
		err := r.Db.Where("gid = ? AND id > ?", gid, 0).Limit(100).Find(&logs).Error
		if err != nil {
			return &output, err
		}
		student.LogUpdated = false
		r.Db.Save(&student)
		output.Override = true
	} else {
		output.Override = false
		if startDate == "" {
			// Get logs by startId and numLog
			err := r.Db.Where("gid = ? AND id > ?", gid, startID).Limit(numLog).Find(&logs).Error
			if err != nil {
				return &output, err
			}
		} else {
			t1, _ := time.Parse(time.RFC3339, startDate)
			err := r.Db.Where("gid = ? AND create_at >= ?", gid, t1).Limit(numLog).Find(&logs).Error
			if err != nil {
				return &output, err
			}
		}
	}

	output.Logs = logs
	return &output, nil
}

// Mutation

func (r *mutationResolver) DelTaskLog(ctx context.Context, logID int) (bool, error) {
	// Create scoped clean db interface
	log := &model.TaskLog{}
	if err := r.Db.Where("id=?", logID).First(log).Error; err != nil {
		return false, err
	}
	var student model.Student
	if err := r.Db.Where("gid = ?", log.Gid).First(&student).Error; err == nil {
		r.Db.Model(&student).Update("log_updated", true)
	}
	if err := r.Db.Delete(log).Error; err != nil {
		return false, err
	}
	return true, nil
}

func updateReviseForward(r *mutationResolver, kpSts *model.KpState, ok bool) error {
	ReviseM := model.ReviseMap{}
	r.Db.Where("gid=?", kpSts.Gid).First(&ReviseM)
	// Create or update the Map
	ReviseM.Gid = kpSts.Gid
	ReviseM.Name = kpSts.Name
	tmpRUnit := []model.ReviseUnit{}
	tmpUnit := model.ReviseUnit{}
	json.Unmarshal(ReviseM.ReviseList, &tmpRUnit)
	for _, item := range tmpRUnit {
		if item.Kp == kpSts.Kp {
			tmpUnit = item
			break
		}
	}
	if tmpUnit.Kp != kpSts.Kp {
		tmpUnit.Kp = kpSts.Kp
		tmpUnit.StartAt = time.Now()
		tmpUnit.ReviseState = 0
		tmpUnit.ReviseCount = 0
	} else {
		tmpUnit.ReviseCount++
		tmpUnit.LastRevisDate = time.Now()
	}
	if ok == true {
		if tmpUnit.ReviseState < 6 {
			tmpUnit.ReviseState++
			tmpUnit.DueDate = time.Now().AddDate(0, 0, model.ReviseDay[tmpUnit.ReviseState])
		} else {
			// Revise complete
			tmpUnit.ReivseComplete = true
			tmpUnit.CompleteAt = time.Now()
			// Set Master level = 3, and we will not set this kp as assignment anymore
			kpSts.Master = 3
		}
	} else {
		tmpUnit.ReviseState = 0
		tmpUnit.DueDate = time.Now().AddDate(0, 0, model.ReviseDay[tmpUnit.ReviseState])
	}

	addNew := true
	for _, item := range tmpRUnit {
		if item.Kp == kpSts.Kp {
			item = tmpUnit
			addNew = false
			break
		}
	}
	if addNew == true {
		tmpRUnit = append(tmpRUnit, tmpUnit)
	}
	ReviseM.ReviseList, _ = json.Marshal(tmpRUnit)
	r.Db.Save(&ReviseM)

	return nil
}
