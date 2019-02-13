package test

import (
	"context"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/stretchr/testify/assert"
	"gitlab.com/dhq_server/config/db"
	"gitlab.com/dhq_server/data/entity"
	"gitlab.com/dhq_server/data/repository_impl"
)

var Token = "tokenTest"
var DB = db.MockSql()

var Todo = entity.Todo{
	Ts:   123,
	Name: "Test",
	JobSite: entity.JobSite{
		Ts:         123,
		Identifier: "Test",
		Name:       "Test",
		BranchID:   1,
	},
}

func created(_db *gorm.DB, entity *entity.Todo) error {
	if err := _db.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func TestGetTodo(t *testing.T) {

	a := repository_impl.NewTodoRepository(DB)
	todos, err := a.Get(context.TODO(), Token)

	if err != nil {
		print(err.Error())
		if Token == "" && err.Error() != "An authorization header is required" {
			if Token != "" && err.Error() != "Unauthorized" {
				t.Errorf("An error has arisen in the petition %s", err.Error())
			}
		}
	}
	if len(todos) == 0 {
		if err := created(DB, &Todo); err != nil {
			t.Errorf("An error has arisen in the petition %s", err)
		}
	} else {
		if err := DB.Delete(todos[0]).Error; err != nil {
			t.Errorf("An error has arisen in the petition %s", err)
		}
	}
}
func TestGetAllTodo(t *testing.T) {

	a := repository_impl.NewTodoRepository(DB)
	todos, err := a.GetAll(context.TODO(), Token)

	if err != nil {
		print(err.Error())
		if Token == "" && err.Error() != "An authorization header is required" {
			if Token != "" && err.Error() != "Unauthorized" {
				t.Errorf("An error has arisen in the petition %s", err.Error())
			}
		}
	}
	if len(todos) == 0 {
		if err := created(DB, &Todo); err != nil {
			t.Errorf("An error has arisen in the petition %s", err)
		}
	} else {
		if err := DB.Delete(todos[0]).Error; err != nil {
			t.Errorf("An error has arisen in the petition %s", err)
		}
	}
}
func TestSyncTodo(t *testing.T) {

	a := repository_impl.NewTodoRepository(DB)
	todos, err := a.Sync(context.TODO(), 0, Token)

	if err != nil {
		print(err.Error())
		if Token == "" && err.Error() != "An authorization header is required" {
			if Token != "" && err.Error() != "Unauthorized" {
				t.Errorf("An error has arisen in the petition %s", err.Error())
			}
		}
	}
	if len(todos) == 0 {
		if err := created(DB, &Todo); err != nil {
			t.Errorf("An error has arisen in the petition %s", err)
		}
	}
}

// func TestGet(t *testing.T) {

// 	DB := db.MockSql()

// 	todo := entity.Todo{
// 		Name:    "Jinzhu",
// 		JobSite: {},
// 		Ts:      123,
// 	}

// 	DB.Create(&todo)

// 	//ctx := new(routing.Context)
// 	a := repository_impl.NewTodoRepository(DB)
// 	result, err := a.Get(context.TODO(), "532eaabd9574880dbf76b9b8cc00832c20a6ec113d682299550d7a6e0f345e25")
// 	assert.NoError(t, err)
// 	if len(result) != 2 {
// 		t.Errorf("Returned sets is not equal to 1. Received %d", len(result))
// 	}
// 	if result[0].ID != 1 {
// 		t.Errorf("ID is not equal. Got %v", result[0].ID)
// 	}
// }
