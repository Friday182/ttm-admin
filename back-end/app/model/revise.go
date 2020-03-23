package model

import (
	"time"
)

// Revise state as index, value is the days for the next due date
var ReviseDay = [...]int{1, 2, 5, 10, 15, 20, 30}

// Each student has unique revise map, it is plan of when and which kp will be repeated
// And move to the master level when complete
type ReviseUnit struct {
	Kp             string
	StartAt        time.Time
	CompleteAt     time.Time
	DueDate        time.Time
	LastRevisDate  time.Time
	ReviseState    int // predefined revise state
	ReivseComplete bool
	ReviseCount    int
}

type ReviseMap struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	Name       string
	Gid        string `gorm:"index;unique_index"`
	ReviseList []byte // ReviseUnit list
}
