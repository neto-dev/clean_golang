package repository

/*
Importamos las librerias

We import libraries
*/
import (
	"context"

	"github.com/clean_golang/utils"

	"github.com/clean_golang/data/entity"
)

/*
Declaramos en la siguiente interfaz los metodos correspondiente al repositotio de la entidad declarada

We declare in the following interface the methods corresponding to the repositotio of the declared entity
*/
type TodoRepository interface {
	Get(ctx context.Context) ([]*entity.Todo, error)
	GetAll(ctx context.Context) ([]*entity.Todo, error)
	Filters(ctx context.Context, _filter *utils.Filters) ([]*entity.Todo, error)
	GetByID(ctx context.Context, _id uint) (*entity.Todo, error)
	Create(ctx context.Context, _area *entity.Todo) (uint, error)
	Update(ctx context.Context, _area *entity.Todo) (uint, error)
	Delete(ctx context.Context, _area *entity.Todo) (bool, error)
}
