package test

import (
	"context"
	"testing"

	"github.com/clean_golang/data/entity"
	usecase "github.com/clean_golang/domain/usecase_impl"
	"github.com/clean_golang/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetArea(t *testing.T) {
	mockAreaRepo := new(mocks.AreaRepository)
	mockArea := &entity.Area{
		Ts:        123,
		JobSiteID: 1,
		Name:      "Area 1",
	}

	mockListArea := make([]*entity.Area, 0)
	mockListArea = append(mockListArea, mockArea)
	mockAreaRepo.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(mockListArea, nil)

	u := usecase.NewAreaUsecase(mockAreaRepo)
	//ctx := new(routing.Context)
	list, err := u.Get(context.TODO(), "tokenTest")
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArea))

	mockAreaRepo.AssertExpectations(t)
}

func TestGetAllArea(t *testing.T) {
	mockAreaRepo := new(mocks.AreaRepository)
	mockArea := &entity.Area{
		Ts:        123,
		JobSiteID: 1,
		Name:      "Area 1",
	}

	mockListArea := make([]*entity.Area, 0)
	mockListArea = append(mockListArea, mockArea)
	mockAreaRepo.On("GetAll", mock.Anything, mock.AnythingOfType("string")).Return(mockListArea, nil)

	u := usecase.NewAreaUsecase(mockAreaRepo)
	//ctx := new(routing.Context)
	list, err := u.GetAll(context.TODO(), "tokenTest")
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArea))

	mockAreaRepo.AssertExpectations(t)
}

func TestSyncArea(t *testing.T) {
	mockAreaRepo := new(mocks.AreaRepository)
	mockArea := &entity.Area{
		Ts:        123,
		JobSiteID: 1,
		Name:      "Area 1",
	}

	mockListArea := make([]*entity.Area, 0)
	mockListArea = append(mockListArea, mockArea)
	mockAreaRepo.On("Sync", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(mockListArea, nil)

	u := usecase.NewAreaUsecase(mockAreaRepo)
	//ctx := new(routing.Context)
	list, err := u.Sync(context.TODO(), 0, "tokenTest")
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArea))

	mockAreaRepo.AssertExpectations(t)
}
