package model

import (
// "errors"
// "github.com/gogf/gf/util/grand"
// "time"
)

type KpDescription struct {
	ID          int `gorm:"primary_key"`
	Kp          string
	Title       string
	Description string
	Review      bool
	Note        string
}

type Question struct {
	ID           int `gorm:"primary_key"`
	QueIdx       int
	Kp           string `sql:"index"`
	StdSec       int
	AnswerType   string
	QuestionType string
	UpTexts      []byte
	DownTexts    []byte
	Formula      []byte
	Options      []byte
	Answers      []byte
	Charts       []byte
	Clocks       []byte
	Tables       []byte
	Shapes       []byte
	Helper       bool
	Imgs         []byte
	Tags         []byte		// for quiz
	AnswerText   string		// for quiz to show the real answer
	Tips         []byte		// Not used yet
	Status       int			// 0 = not finish, 1 = need review, 2 = complete
}
