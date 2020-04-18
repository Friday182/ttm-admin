package model

import (
	"errors"
	//ws "github.com/friday182/ttm-admin/app/api/websocket"
	"time"

	"github.com/gogf/gf/util/grand"
)

type Student struct {
	ID              int    `gorm:"primary_key"`
	Name            string `json:"Name"`
	Password        string `json:"Password"`
	Gid             string `gorm:"index;unique_index"`
	MentorEmail     string `json:"MentorEmail"`
	Gender          string
	Age             int `json:"Age"`
	Country         string
	City            string
	School          string
	Activity        int
	Stickers        int
	Stars           int
	TaskDoneCount   int
	TaskPassCount   int
	LogUpdated      bool
	IsMember        bool
	MembershipStart time.Time
	MembershipStop  time.Time
	LastLoginTime   time.Time
	Quiz            []byte // array of QuizUnit
	StickerLog      []byte
	Contacts        []byte
	SubjectState    []byte // array of subject status
	KpPlan          []byte // array of KpUnit
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

type KpAssignment struct {
	Kp     string
	QueIdx int
	NumQue int
	Done   bool
	Pass   bool
	Desc   string
}

type QuizUnit struct {
	QuizId  string
	Done    bool
	Pass    bool
	Desc    string
	Percent int
}

// This is log of subject status for a student
type Subject struct {
	Name         string
	IsEnabled    bool
	AiEnabled    bool
	Level        int
	Assignment   []KpAssignment
	LastDoneTime time.Time
}

type KpUnit struct {
	Kp     string
	Weight int // Gigger then Higher priority, so we always need a sort array of it
}

type KpUnitMethod interface {
	GetWeight() int
}

func (u *KpUnit) GetWeight() int {
	return u.Weight
}

func AddStudent(new *Student) (*Student, error) {
	u := Student{}
	if new.MentorEmail != "" && len(new.Name) > 1 && new.Age > 4 && new.Age < 12 {
		if ttmDb.Where(Student{MentorEmail: new.MentorEmail, Name: new.Name, Age: new.Age}).Find(&u).RecordNotFound() {
			var tmp string
			if len(new.Name) > 3 {
				tmp = new.Name[:3]
			} else {
				tmp = new.Name
			}
			new.Password = tmp + grand.Digits(3)
			err := ttmDb.Save(new).Error
			return new, err
		}
		return &u, errors.New("Duplicated Student")
	}
	return &u, errors.New("Invalid Parameters")
}

func DelStudent(gid string) error {
	u := Student{}
	ttmDb := ttmDb.New().Begin()

	if ttmDb.Where(Student{Gid: gid}).Find(&u).RecordNotFound() {
		return errors.New("Not Found Student")
	}
	if err := ttmDb.Delete(u).Error; err != nil {
		ttmDb.RollbackUnlessCommitted()
		return ttmDb.Error
	}
	ttmDb = ttmDb.Commit()
	return nil
}

func ValidStudent(user string, pw string) bool {
	u := Student{}
	if ttmDb.Where(Student{Name: user, Password: pw}).Find(&u).RecordNotFound() {
		return false
	}
	return true
}
