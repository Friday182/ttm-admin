package resolvers_test

import (
	"context"
	"testing"

	//"time"

	//"github.com/friday182/ttm-admin/app/model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestGetQuizLogs(t *testing.T) {
	var c context.Context
	list, err := q.Query().GetQuizLog(c, "test")
	if len(list) != 0 && err != nil {
		t.Fatal("Get task log failed", err)
	}
}
