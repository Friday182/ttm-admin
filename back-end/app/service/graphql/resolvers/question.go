package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/friday182/ttm-go/app/model"
	"github.com/friday182/ttm-go/app/service/graphql/generated"
	"github.com/jinzhu/gorm"
)

// Query
func (r *queryResolver) GetKpDescripitions(ctx context.Context) ([]*model.KpDescription, error) {
	desList := []*model.KpDescription{}
	err := r.Qdb.Find(&desList).Error
	return desList, err
}

func (r *queryResolver) GetQuestions(ctx context.Context, gid string, kp string) ([]*model.Question, error) {
	var user model.User
	queList := []*model.Question{}
	err := r.Db.Where("gid = ?", gid).Find(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return queList, err
	}

	err = r.Qdb.Where("kp=?", kp).Limit(50).Find(&queList).Error
	return queList, err
}

// Mutation
func (r *mutationResolver) AddQuestion(ctx context.Context, que generated.AddQuestionInput) (bool, error) {
	panic("not implemented")
}
