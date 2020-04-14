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

var answerOptions []string = []string{"1", "2", "3", "4", "5"}

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

	err := r.Db.Where("gid = ?", que.Gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, err
	}

	newQue.QueIdx = que.QueIdx
	newQue.Kp = que.Kp
	newQue.StdSec = que.StdSec
	newQue.AnswerType = que.AnswerType
	newQue.QuestionType = que.QuestionType
	newQue.Imgs, _ = json.Marshal([]string{""})
	newQue.Tips, _ = json.Marshal([]string{""})
	newQue.Helper = false

	tmpStr := strings.TrimSuffix(que.UpTexts, "\n")
	tmpArray := strings.Split(tmpStr, "\n")
	tmpArray = remove_empty(tmpArray)
	tmp, err := json.Marshal(tmpArray)
	newQue.UpTexts = tmp

	tmpStr = strings.TrimSuffix(que.DownTexts, "\n")
	tmpArray = strings.Split(tmpStr, "\n")
	tmpArray = remove_empty(tmpArray)
	tmp, err = json.Marshal(tmpArray)
	newQue.DownTexts = tmp

	tmpArray = strings.Split(que.Formula, "\n")
	tmp, err = json.Marshal(tmpArray)
	newQue.Formula = tmp

	tmpArray = strings.Split(que.Options, "||")
	//tmpArray = remove_empty(tmpArray)
	tmp, err = json.Marshal(tmpArray)
	newQue.Options = tmp

	tmpArray = strings.Split(que.Charts, "||")
	//tmpArray = remove_empty(tmpArray)
	tmp, err = json.Marshal(tmpArray)
	newQue.Charts = tmp

	tmpStr = strings.TrimSuffix(que.Tables, "||")
	tmpArray = strings.Split(tmpStr, "||")
	//tmpArray = remove_empty(tmpArray)
	tmp, err = json.Marshal(tmpArray)
	newQue.Tables = tmp

	tmpStr = strings.TrimSuffix(que.Shapes, "||")
	tmpArray = strings.Split(tmpStr, "||")
	tmpArray = remove_empty(tmpArray)
	//temp = []string{}
	//err = json.Unmarshal([]byte(tmpStr), &temp)
	tmp, err = json.Marshal(tmpArray)
	newQue.Shapes = tmp

	tmpArray = strings.Split(que.Clocks, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Clocks = tmp

	tmpStr = strings.TrimSuffix(que.Tags, "||")
	tmpArray = strings.Split(tmpStr, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Tags = tmp

	// Check Answer is valid
	tmpArray = strings.Split(que.Answers, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Answers = tmp

	tmpQue := model.Question{}
	err = r.Qdb.Where("kp = ? AND que_idx = ?", que.Kp, que.QueIdx).Find(&tmpQue).Error
	if gorm.IsRecordNotFoundError(err) {
		r.Qdb.Save(&newQue)
		return true, nil
	} else {
		newQue.ID = tmpQue.ID
		r.Qdb.Save(&newQue)
		return true, nil
	}
}

func remove_empty (s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
