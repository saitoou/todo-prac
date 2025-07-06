package repository

// TODO Test

// import (
// 	"context"
// 	"regexp"
// 	"testing"
// 	"time"
// 	"todo-golang/domain/entity"
// 	"todo-golang/testutils"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/golang/mock/gomock"
// )

// func Test_TodoRepository_FindAll(t *testing.T) {
// 	ctx := context.Background()
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockdb, gormMock := testutils.Mock(t)
// 	repo := NewTodoRepository(mockdb)
// 	now := time.Now().UTC().Truncate(time.Millisecond)

// 	tests := []struct {
// 		name     string
// 		mockFunc func()
// 		want     []*entity.Todo
// 		wantErr  error
// 	}{
// 		{
// 			name: "When Success",
// 			mockFunc: func() {
// 				rows := sqlmock.NewRows([]string{"id", "title", "content", "user_id", "created_at"})
// 				rows.AddRow(1, "title1", "title1", 1, now)
// 				rows.AddRow(2, "title2", "title2", 2, now)
// 				rows.AddRow(3, "title3", "title3", 3, now)

// 				gormMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM postgres.todo WHERE "))

// 			},
// 		},
// 	}

// }
