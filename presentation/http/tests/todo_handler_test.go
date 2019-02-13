package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/clean_golang/data/entity"
	"github.com/clean_golang/mocks"
	handlerHttp "github.com/clean_golang/presentation/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTodo(t *testing.T) {
	mockTodo := &entity.Todo{
		Name:        "Todo 1",
		Description: "Todo Description",
	}

	mockUCase := new(mocks.TodoUsecase)

	mockListTodo := make([]*entity.Todo, 0)
	mockListTodo = append(mockListTodo, mockTodo)

	mockUCase.On("Get", mock.Anything).Return(mockListTodo, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/v2/todos/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	handler := handlerHttp.HttpTodoHandler{
		TodoUsecase: mockUCase,
	}
	handler.Get(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestGetAllTodo(t *testing.T) {
	mockTodo := &entity.Todo{
		Name:        "Todo 1",
		Description: "Todo Description",
	}

	mockUCase := new(mocks.TodoUsecase)

	mockListTodo := make([]*entity.Todo, 0)
	mockListTodo = append(mockListTodo, mockTodo)

	mockUCase.On("GetAll", mock.Anything).Return(mockListTodo, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/v2/todos/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	handler := handlerHttp.HttpTodoHandler{
		TodoUsecase: mockUCase,
	}
	handler.GetAll(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}
