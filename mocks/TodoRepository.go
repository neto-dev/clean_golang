package mocks

import (
	"context"

	"github.com/clean_golang/data/entity"
	"github.com/clean_golang/utils"
	mock "github.com/stretchr/testify/mock"
)

type TodoRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *TodoRepository) Get(ctx context.Context) ([]*entity.Todo, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Todo); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TodoRepository) GetAll(ctx context.Context) ([]*entity.Todo, error) {

	ret := _m.Called(ctx)

	var r0 []*entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Todo); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TodoRepository) GetByID(ctx context.Context, id uint) (*entity.Todo, error) {

	ret := _m.Called(ctx)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entity.Todo); ok {
		r0 = rf(ctx, id, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TodoRepository) Filters(ctx context.Context, filter *utils.Filters) ([]*entity.Todo, error) {

	return nil, nil
}

func (_m *TodoRepository) Create(ctx context.Context, todo *entity.Todo) (uint, error) {

	return 0, nil
}

func (_m *TodoRepository) Update(ctx context.Context, todo *entity.Todo) (uint, error) {

	return 0, nil
}

func (_m *TodoRepository) Delete(ctx context.Context, todo *entity.Todo) (bool, error) {

	return false, nil
}
