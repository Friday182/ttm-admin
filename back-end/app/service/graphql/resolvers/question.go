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
	err := r.QuizDb.Find(&desList).Error
	return desList, err
}

func (r *queryResolver) GetQuestions(ctx context.Context, gid string, kp string) ([]*model.Question, error) {
	var user model.User
	queList := []*model.Question{}
	err := r.TtmDb.Where("gid = ?", gid).Find(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return queList, err
	}

	err = r.QuizDb.Where("kp=?", kp).Order("que_idx").Limit(50).Find(&queList).Error
	return queList, err
}

// Mutation
func (r *mutationResolver) AddQuestion(ctx context.Context, que generated.AddQuestionInput) (*model.Question, error) {
	newQue := model.Question{}
	var user model.User

	err := r.TtmDb.Where("gid = ?", que.Gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return &newQue, err
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
	tmpArray = removeEmpty(tmpArray)
	tmp, err := json.Marshal(tmpArray)
	newQue.UpTexts = tmp

	tmpStr = strings.TrimSuffix(que.DownTexts, "\n")
	tmpArray = strings.Split(tmpStr, "\n")
	tmpArray = removeEmpty(tmpArray)
	tmp, err = json.Marshal(tmpArray)
	newQue.DownTexts = tmp

	// temp add process for English text question
	if strings.Contains(que.Kp, "EQ_") {
		newQue.DownTexts = processEnText(tmpArray)
	}

	tmpArray = strings.Split(que.Formula, "\n")
	tmp, err = json.Marshal(tmpArray)
	newQue.Formula = tmp

	tmpArray = strings.Split(que.Options, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Options = tmp

	tmpStr = strings.TrimSuffix(que.Charts, "||")
	tmpArray = strings.Split(tmpStr, "||")
	tmpArray = removeEmpty(tmpArray)
	tmp, err = json.Marshal(tmpArray)
	newQue.Charts = tmp

	tmpStr = strings.TrimSuffix(que.Tables, "||")
	tmpArray = strings.Split(tmpStr, "||")
	//tmpArray = removeEmpty(tmpArray)
	tmp, err = json.Marshal(tmpArray)
	newQue.Tables = tmp

	tmpStr = strings.TrimSuffix(que.Shapes, "||")
	tmpArray = strings.Split(tmpStr, "||")
	tmpArray = removeEmpty(tmpArray)
	//temp = []string{}
	//err = json.Unmarshal([]byte(tmpStr), &temp)
	tmp, err = json.Marshal(tmpArray)
	newQue.Shapes = tmp

	tmpArray = strings.Split(que.Clocks, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Clocks = tmp

	tmpStr = strings.TrimSuffix(que.Tags, "||")
	tmpTags := strings.Split(tmpStr, "||")
	tmp, err = json.Marshal(tmpTags)
	newQue.Tags = tmp

	// Check Answer is valid
	tmpArray = strings.Split(que.Answers, "||")
	tmp, err = json.Marshal(tmpArray)
	newQue.Answers = tmp

	tmpQue := model.Question{}
	err = r.QuizDb.Where("kp = ? AND que_idx = ?", que.Kp, que.QueIdx).Find(&tmpQue).Error
	if gorm.IsRecordNotFoundError(err) {
		err = r.QuizDb.Save(&newQue).Error
	} else {
		newQue.ID = tmpQue.ID
		err = r.QuizDb.Save(&newQue).Error
	}

	quiz := model.Quiz{}
	err = r.QuizDb.Where("quiz_id = ?", que.Kp).Find(&quiz).Error
	if err == nil {
		if quiz.Status != "WIP" {
			quiz.Status = "WIP"
		}
		// update quiz details
		tmpDetails := model.QuizDetails{}
		err = json.Unmarshal(quiz.Details, &tmpDetails)
		for i := 0; i < len(tmpTags); i++ {
			if strings.Contains(tmpTags[i], "MA") {
				tmpDetails.Ma++
			} else if strings.Contains(tmpTags[i], "MB") {
				tmpDetails.Mb++
			} else if strings.Contains(tmpTags[i], "MC") {
				tmpDetails.Mc++
			} else if strings.Contains(tmpTags[i], "MD") {
				tmpDetails.Md++
			} else if strings.Contains(tmpTags[i], "ME") {
				tmpDetails.Me++
			} else if strings.Contains(tmpTags[i], "MF") {
				tmpDetails.Mf++
			} else if strings.Contains(tmpTags[i], "MG") {
				tmpDetails.Mg++
			} else if strings.Contains(tmpTags[i], "MH") {
				tmpDetails.Mh++
			} else if strings.Contains(tmpTags[i], "MI") {
				tmpDetails.Mi++
			} else if strings.Contains(tmpTags[i], "MJ") {
				tmpDetails.Mj++
			}
		}
		quiz.Details, _ = json.Marshal(tmpDetails)
		r.QuizDb.Save(&quiz)
	}
	return &newQue, nil
}

func removeEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func processEnText(t []string) []byte {
	strs := []string{}
	for _, item := range t {
		str := strings.Split(item, " ")
		strs = append(strs, str...)
	}
	opt1 := checkPuncRule1(strs)
	opt2 := checkPuncRule2(strs)

	tmp, _ := json.Marshal(t)
	return tmp
}
/*
func processEnText(test []string) []byte {
	textList := [][]string{}
	b := []byte(" ")
	for _, item := range test {
		head := 0
		tail := 0
		for head = 0; head < len(item); head++ {
			if head-tail > 50 && item[head] == b[0] {
				str := item[tail:head]
				strs := strings.Split(str, " ")
				strs = removeEmpty(strs)
				textList = append(textList, strs)
				tail = head
			}
		}
		if tail != head {
			str := item[tail:head]
			strs := strings.Split(str, " ")
			strs = removeEmpty(strs)
			textList = append(textList, strs)
			tail = head
		}
	}
	tmp, _ := json.Marshal(textList)
	return tmp
}
*/

// Punctuation rule 1, Capital letter
func checkPuncRule1 (s []string) []int {
	tmp := []int{}
	for idx, item := range s {
		if item[0] > 0x30 {
			tmp = append(tmp, idx)
		}
	}

	return tmp
}
