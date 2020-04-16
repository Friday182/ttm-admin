package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/friday182/ttm-admin/app/model"
	"github.com/friday182/ttm-admin/app/service/graphql/generated"
	"github.com/gogf/gf/util/grand"
	"github.com/jinzhu/gorm"
)

// Query
func (r *queryResolver) GetUsers(ctx context.Context, role string) ([]*model.User, error) {
	userList := []*model.User{}
	err := r.Db.Where("role = ?", role).Find(&userList).Error
	return userList, err
}

func (r *mutationResolver) AddUser(ctx context.Context, user generated.AddUserInput) (bool, error) {
	newUser := model.User{}

	for {
		user.Username = user.Name + grand.Digits(3)
		user.Password = grand.Digits(4)
		err := r.Db.Where("username = ?", user.Username).Find(&newUser).Error
		if gorm.IsRecordNotFoundError(err) {
			newUser.Username = user.Username
			newUser.Password = user.Password
			newUser.Name = user.Name
			newUser.Email = user.Email
			newUser.Role = user.Role
			newUser.Mobile = user.Mobile
			newUser.Comment = user.Comment
			newUser.Gid = strconv.FormatInt(time.Now().UnixNano(), 36)
			newUser.LastLoginTime = time.Now()

			err = r.Db.Save(&newUser).Error

			return true, err
		}
	}
}

func (r *mutationResolver) DelUser(ctx context.Context, gid string) (bool, error) {
	var user model.User
	var result bool = false
	err := r.Db.Where("gid=?", gid).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		result = false
	} else {
		if err = r.Db.Delete(&user).Error; err != nil {
			result = false
		} else {
			result = true
		}
	}

	return result, err
}

func (r *mutationResolver) UserLogin(ctx context.Context, username string, password string) (*model.User, error) {
	var user model.User
	if r.Db.Where("username = ? AND password = ?", username, password).Find(&user).RecordNotFound() {
		return &user, errors.New("failed")
	} else {
		user.LastLoginTime = time.Now()
		r.Db.Save(&user)

		return &user, nil
	}
}
