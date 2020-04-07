package resolvers_test

import (
	"testing"
	//"encoding/json"
	//"time"
	//"github.com/friday182/ttm-admin/app/service/graphql/resolvers"
	//"github.com/friday182/ttm-admin/app/model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"github.com/gogf/gf/util/grand"
)

func TestGetQuestions(t *testing.T) {
	queList, err := q.Query().GetQuestions(c, studentGid, 0, "MA1", 10)
	if err != nil {
		t.Fatal("Get question failed", err)
	}
	if len(queList) == 0 {
		t.Fatal("Get question failed")
	}
}
