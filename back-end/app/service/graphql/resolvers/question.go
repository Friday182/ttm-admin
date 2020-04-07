package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/friday182/ttm-admin/app/model"
	"github.com/friday182/ttm-admin/app/service/graphql/generated"
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
	newQue := model.Question{}
	var user model.User

	err := r.Db.Where("gid = ?", que.Gid).Find(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, err
	}

	newQue.QueIdx = que.QueIdx
	newQue.Kp = que.Kp
	newQue.StdSec = que.StdSec
	newQue.AnswerType = que.AnswerType
	newQue.QuestionType = que.QuestionType
	newQue.Imgs = []byte{}
	newQue.Tips = []byte{}
	newQue.Helper = false

	tmpArray := strings.Split(que.UpTexts, "\n")
	tmp, err := json.Marshal(tmpArray)
	newQue.UpTexts = tmp

	tmpArray = strings.Split(que.DownTexts, "\n")
	tmp, err = json.Marshal(tmpArray)
	newQue.DownTexts = tmp

	tmpArray = strings.Split(que.Formula, "\n")
	tmp, err = json.Marshal(tmpArray)
	newQue.Formula = tmp

	tmpArray = strings.Split(que.Options, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Options = tmp

	tmpArray = strings.Split(que.Charts, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Charts = tmp

	tmpArray = strings.Split(que.Tables, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Tables = tmp

	tmpArray = strings.Split(que.Shapes, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Shapes = tmp

	tmpArray = strings.Split(que.Clocks, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Clocks = tmp

	err = r.Qdb.Where("kp = ? AND que_idx = ?", que.Kp, que.QueIdx).Find(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		r.Qdb.Save(&newQue)
		return true, err
	} else {
		r.Qdb.Save(&newQue)
		return true, err
	}
}
