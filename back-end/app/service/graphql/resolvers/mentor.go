package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	//"strings"
	"github.com/jinzhu/gorm"
	//"github.com/friday182/ttm-admin/app/service/graphql/generated"
	"github.com/friday182/ttm-admin/app/model"
)

// Query
func (r *queryResolver) Mentors(ctx context.Context, email string) ([]*model.Mentor, error) {
	mentorList := []*model.Mentor{}
	err := r.TtmDb.Limit(100).Find(&mentorList).Error
	return mentorList, err
}

// Mutations
func (r *mutationResolver) DelMentor(ctx context.Context, email string) (bool, error) {
	var mentor model.Mentor
	var result bool = false
	// Todo: verify the operation right before do it
	err := r.TtmDb.Where("email=?", email).First(&mentor).Error
	if gorm.IsRecordNotFoundError(err) {
		result = false
	} else {
		if err = r.TtmDb.Delete(&mentor).Error; err != nil {
			result = false
		} else {
			result = true
			r.TtmDb.Commit()
		}
	}

	return result, err
}
