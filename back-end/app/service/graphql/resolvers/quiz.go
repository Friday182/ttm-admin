package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"encoding/json"
	"errors"

	//"time"
	"github.com/friday182/ttm-admin/app/model"
	"github.com/friday182/ttm-admin/app/service/graphql/generated"
	"github.com/jinzhu/gorm"
	//"github.com/gogf/gf/util/gconv"
)

// Query
func (r *queryResolver) GetQuizLog(ctx context.Context, gid string) ([]*model.QuizLog, error) {
	var logs []*model.QuizLog
	err := r.TtmDb.Where("gid = ?", gid).Find(&logs).Error
	return logs, err
}

func (r *queryResolver) QuizLog(ctx context.Context, gid string, quizIdx *string) (*model.QuizLog, error) {
	var log model.QuizLog
	err := r.TtmDb.Where("gid = ? AND quiz_idx=?", gid, quizIdx).Find(&log).Error
	return &log, err
}

func (r *queryResolver) GetQuizReport(ctx context.Context, gid string) (*model.QuizReport, error) {
	var log model.QuizReport
	err := r.TtmDb.Where("gid = ?", gid).Find(&log).Error
	return &log, err
}

func (r *queryResolver) GetQuiz(ctx context.Context, gid string) ([]*model.Quiz, error) {
	var logs []*model.Quiz
	user := model.User{}
	err := r.TtmDb.Where("gid = ?", gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return logs, err
	}
	if user.Role == "operator" {
		err = r.QuizDb.Where("operator=?", user.Username).Find(&logs).Error
	} else {
		err = r.QuizDb.Order("id desc").Limit(100).Find(&logs).Error
	}

	return logs, err
}

// Mutation
func (r *mutationResolver) AddQuiz(ctx context.Context, quiz generated.AddQuizInput) (bool, error) {
	newQuiz := model.Quiz{}
	err := r.QuizDb.Where("quiz_id = ?", quiz.QuizID).Find(&newQuiz).Error
	if gorm.IsRecordNotFoundError(err) {
		tmpQuizDetails := model.QuizDetails{
			Ma: 0, Mb: 0, Mc: 0, Md: 0, Me: 0, Mf: 0, Mg: 0, Mh: 0, Mi: 0, Mj: 0,
		}
		newQuiz.QuizId = quiz.QuizID
		newQuiz.Grade = quiz.Grade
		newQuiz.UsedCount = 0
		newQuiz.Desc = quiz.Desc
		newQuiz.Comment = quiz.Comment
		newQuiz.Status = "NA"
		newQuiz.Operator = quiz.Operator
		newQuiz.Approver = "NA"
		newQuiz.Details, _ = json.Marshal(tmpQuizDetails)
		err = r.QuizDb.Save(&newQuiz).Error

		return true, err
	}
	return false, errors.New("Quiz already exist")
}

func (r *mutationResolver) DelQuiz(ctx context.Context, gid string, quizID string) (bool, error) {
	q := model.Quiz{}
	user := model.User{}

	err := r.TtmDb.Where("gid = ?", gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, err
	}
	if user.Role != "admin" && user.Role != "staff" {
		return false, errors.New("No right to delete quiz")
	}
	err = r.QuizDb.Where("quiz_id=?", quizID).First(&q).Error
	if err == nil {
		err = r.QuizDb.Delete(&q).Error
	}
	return false, err
}

func (r *mutationResolver) UpdateQuiz(ctx context.Context, gid string, quizID string, status string) (bool, error) {
	q := model.Quiz{}
	user := model.User{}

	err := r.TtmDb.Where("gid = ?", gid).First(&user).Error
	if err != nil {
		return false, err
	}
	err = r.QuizDb.Where("quiz_id=?", quizID).First(&q).Error
	if err != nil {
		return false, err
	}

	if status == "Delete" {
		if user.Role != "admin" && user.Role != "staff" {
			return false, errors.New("No right to delete quiz")
		}
		if q.Status == "NA" {
			if err = r.QuizDb.Delete(&q).Error; err == nil {
				return true, nil
			}
		}
		return false, err
	} else if status == "Finish" {
		if q.Operator == user.Username {
			q.Status = "Finish"
			err = r.QuizDb.Save(&q).Error
			return true, err
		}
		return false, errors.New("No valid operator")
	} else if status == "Ready" {
		if user.Role == "operator" {
			return false, errors.New("No right to approval")
		}
		if q.Status != "Ready" {
			q.Status = "Ready"
			q.Approver = user.Username
			err = r.QuizDb.Save(&q).Error
		}
		return true, err
	}

	return false, errors.New("No successful")
}

func (r *mutationResolver) FinishQuiz(ctx context.Context, gid string, quizID string) (bool, error) {
	q := model.Quiz{}
	user := model.User{}

	err := r.TtmDb.Where("gid = ?", gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, err
	}
	err = r.QuizDb.Where("quiz_id=?", quizID).First(&q).Error
	if err == nil {
		if q.Operator == user.Username {
			q.Status = "Finish"
			err = r.QuizDb.Save(&q).Error
			return true, err
		} else {
			return false, errors.New("No valid operator")
		}
	}
	return false, err
}

func (r *mutationResolver) ApprovalQuiz(ctx context.Context, gid string, quizID string) (bool, error) {
	q := model.Quiz{}
	user := model.User{}

	err := r.TtmDb.Where("gid = ?", gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, err
	}
	if user.Role == "operator" {
		return false, errors.New("No right to approval")
	}
	err = r.QuizDb.Where("quiz_id=?", quizID).First(&q).Error
	if err == nil {
		if q.Status != "Ready" {
			q.Status = "Ready"
			err = r.QuizDb.Save(&q).Error
			return true, err
		}
	}
	return false, err
}
