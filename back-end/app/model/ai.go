package model

import (
	"time"
)

// Record for each completed task
type TaskLog struct {
	ID           int `gorm:"primary_key"`
	CreatedAt    time.Time
	Gid          string `gorm:"index"`
	Subject      int
	Name         string
	TotalQue     int
	TotalCorrect int
	TotalSec     int
	TotalScore   int
	Kps          string
}

type KpLog struct {
	Kp         string
	NumQue     int
	CorrectQue int
	MaxSeconds int
	MinSeconds int
	AvgSeconds int
}

// Record for each student, each KP, determined by Gid and Kp
type KpState struct {
	ID           int `gorm:"primary_key"`
	CreatedAt    time.Time
	Name         string
	Gid          string
	Kp           string
	TotalQue     int
	StdSec       int
	TotalSec     int
	AvgCorrect   int
	AvgTime      int
	Master       int    // 3 levels, learning, ready, master
	Logs         []byte // list of KpLog
	ResultList   []byte // list of the last 100 questions result
	WrongList    []byte // list of wrong questions Ids
	WrongHisList []byte // list of history wrong questions Ids, will be used in revise plan
}

type KpMap struct {
	Kp      string `gorm:"primary_key"`
	PreKps  []byte
	NextKps []byte
}
