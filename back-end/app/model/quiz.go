package model

import (
	// "errors"
	// "github.com/gogf/gf/util/grand"
	"time"
)

// Quiz Description each Quiz
type Quiz struct {
	ID        int `gorm:"primary_key"`
	QuizId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	UsedCount int
	Grade     int
	Desc			string
	Status    string
	Operator  string
	Approver  string
	Comment   string
	Details   []byte  // list of QuizDetails
}

// QuizLog Record for each time quiz be completed
type QuizLog struct {
	ID          int `gorm:"primary_key"`
	CreatedAt   time.Time
	Name				string
	Gid         string `gorm:"index"`
	QuizId      string // example: MMT1
	TotalQue    int
	DoneQue     int
	CorrectQue  int
	WrongQue    int
	CorrectPer  int
	StdSec      int
	UsedSec     int
	ResultList  []byte // list of QueUnit
	WrongList   []byte // list of QueUnit
	WeightedKps []byte // string array of KPs
	Feedback    string `gorm:"size:1000"`
	Comment     string `gorm:"size:1000"`
}

// QuizReport Create a QuizReport for each student, it will be updated everytime complete of a quiz
type QuizReport struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Gid        string `gorm:"index"`
	TotalQuiz  int
	TotalQue   int
	DoneQue    int
	CorrectQue int
	TotalSec   int
	QuizList   []byte // string array of quiz idx
	Feedback   string `gorm:"size:1000"`
	Comment    string `gorm:"size:1000"`
}

// QuizDetails quiz question statistics
type QuizDetails struct{
	Ma int
	Mb int
	Mc int
	Md int
	Me int
	Mf int
	Mg int
	Mh int
	Mi int
	Mj int
}
