package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"errors"
	//"time"
	"github.com/friday182/ttm-go/app/model"
	"github.com/friday182/ttm-go/app/service/graphql/generated"
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

func (r *queryResolver) GetQuiz(ctx context.Context) ([]*model.Quiz, error) {
	var logs []*model.Quiz
	err := r.Qdb.Limit(100).Find(&logs).Error
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
