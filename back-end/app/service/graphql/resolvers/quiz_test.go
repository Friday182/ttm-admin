package resolvers_test

import (
	"context"
	"testing"
	"encoding/json"
	//"time"
	"github.com/friday182/ttm-go/app/service/graphql/resolvers"
	//"github.com/friday182/ttm-go/app/model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gogf/gf/util/grand"
)

func TestGetQuizLogs(t *testing.T) {
	var c context.Context
	list, err := q.Query().GetQuizLog(c, "test")
	if len(list) != 0 && err != nil {
		t.Fatal("Get task log failed", err)
	}
}
