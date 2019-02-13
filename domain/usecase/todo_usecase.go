package usecase

import (
	"context"

	"github.com/clean_golang/data/entity"
	"github.com/clean_golang/utils"
)

/*
Declaramos en la siguiente interfaz los metodos correspondiente a los casos de usos de la entidad declarada

We declare in the following interface the methods corresponding to the cases of uses of the declared entity
*/
type TodoUsecase interface {
	Get(ctx context.Context) ([]*entity.Todo, error)
	GetAll(ctx context.Context) ([]*entity.Todo, error)
	GetByID(ctx context.Context, _id uint) (*entity.Todo, error)
	Filters(ctx context.Context, _filter *utils.Filters) ([]*entity.Todo, error)
	Create(ctx context.Context, area *entity.Todo) (*entity.Todo, error)
	Update(ctx context.Context, area *entity.Todo) (*entity.Todo, error)
	Delete(ctx context.Context, _id uint) (bool, error)
}
