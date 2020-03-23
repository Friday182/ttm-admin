package resolvers_test

import (
	"context"
	"encoding/json"
	"testing"
	//"time"
	"github.com/friday182/ttm-go/app/service/graphql/resolvers"
	//"github.com/friday182/ttm-go/app/model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestGetTaskLogs(t *testing.T) {
	var c context.Context
	list, err := q.Query().GetTaskLogs(c, "test", "", 0, 10)
	if len(list.Logs) != 0 && err != nil {
		t.Fatal("Get task log failed", err)
	}
}
