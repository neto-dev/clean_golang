package mocks

import (
	"context"

	"github.com/clean_golang/data/entity"
	"github.com/clean_golang/utils"
	mock "github.com/stretchr/testify/mock"
)

type TodoUsecase struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *TodoUsecase) Get(ctx context.Context) ([]*entity.Todo, error) {
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

func (_m *TodoUsecase) GetAll(ctx context.Context) ([]*entity.Todo, error) {

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

func (_m *TodoUsecase) GetByID(ctx context.Context, id uint) (*entity.Todo, error) {

	return nil, nil
}

func (_m *TodoUsecase) Filters(ctx context.Context, filter *utils.Filters) ([]*entity.Todo, error) {

	return nil, nil
}

func (_m *TodoUsecase) Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {

	return nil, nil
}

func (_m *TodoUsecase) Update(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {

	return nil, nil
}

func (_m *TodoUsecase) Delete(ctx context.Context, id uint) (bool, error) {

	return false, nil
}
