package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
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
	err := r.Db.Where("gid = ?", gid).Find(&logs).Error
	return logs, err
}

func (r *queryResolver) QuizLog(ctx context.Context, gid string, quizIdx *string) (*model.QuizLog, error) {
	var log model.QuizLog
	err := r.Db.Where("gid = ? AND quiz_idx=?", gid, quizIdx).Find(&log).Error
	return &log, err
}

func (r *queryResolver) GetQuizReport(ctx context.Context, gid string) (*model.QuizReport, error) {
	var log model.QuizReport
	err := r.Db.Where("gid = ?", gid).Find(&log).Error
	return &log, err
}

func (r *queryResolver) GetQuiz(ctx context.Context, gid string) ([]*model.Quiz, error) {
	var logs []*model.Quiz
	user := model.User{}
	err := r.Db.Where("gid = ?", gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return logs, err
	}
	if user.Role == "operator" {
		err = r.Qdb.Where("operator=?",user.Username).Find(&logs).Error
	} else {
		err = r.Qdb.Order("id desc").Limit(100).Find(&logs).Error
	}

	return logs, err
}

// Mutation
func (r *mutationResolver) AddQuiz(ctx context.Context, quiz generated.AddQuizInput) (bool, error) {
	newQuiz := model.Quiz{}
	err := r.Qdb.Where("quiz_id = ?", quiz.QuizID).Find(&newQuiz).Error
	if gorm.IsRecordNotFoundError(err) {
		newQuiz.QuizId = quiz.QuizID
		newQuiz.Grade = quiz.Grade
		newQuiz.Desc = quiz.Desc
		newQuiz.Comment = quiz.Comment
		newQuiz.Status = "NA"
		newQuiz.Operator = quiz.Operator
		newQuiz.Approver = "NA"
		err = r.Qdb.Save(&newQuiz).Error

		return true, err
	}
	return false, errors.New("Quiz already exist!")
}

func (r *mutationResolver) DelQuiz(ctx context.Context, gid string, quizID string) (bool, error) {
	q := model.Quiz{}
	user := model.User{}

	err := r.Db.Where("gid = ?", gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, err
	}
	if user.Role != "admin" && user.Role != "staff" {
		return false, errors.New("No right to delete quiz")
	}
	err = r.Qdb.Where("quiz_id=?", quizID).First(&q).Error
	if err == nil {
		if err = r.Qdb.Delete(&q).Error; err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}
