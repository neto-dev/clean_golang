package repository_impl

/*
Importamos las librerias

We import libraries
*/

import (
	"context"

	"github.com/clean_golang/data/entity"
	"github.com/clean_golang/data/repository"
	"github.com/clean_golang/utils"
	"github.com/jinzhu/gorm"
)

/*
Estructura la cual almacena la base de datos con la que se va a trabajar en el momento de ejecucion, todas las funciones dependeran de esta para poder acceder a las propiedades de la base de datos

Structure which stores the database with which you will work at the time of execution, all functions depend on this to access the properties of the database
*/
type TodoRepository struct {
	DB *gorm.DB
}

/*
Funcion encargade de crear una nueva instancia del repositorio para poder realizar la inyeccion de dependencias

Function in charge of creating a new instance of the repository to be able to make the injection of dependencies.

*/
func NewTodoRepository(Conn *gorm.DB) repository.TodoRepository {
	/*
		Retornamos la instancia del repositorio

		We return the instance from the repository
	*/
	return &TodoRepository{Conn}
}

/*
Funcion encargada de conectarse con la base de datos y obtener todos los registros de la entidad

Function in charge of connecting to the database and obtaining all the records of the entity.

*/
func (_repo_ *TodoRepository) Get(_ctx context.Context) ([]*entity.Todo, error) {
	/*
	 Se declara la siguiente variable a la cual se le asignara la estructura de la entidad

	  The following variable is declared to which the structure of the entity will be assigned
	*/
	todos := []*entity.Todo{}

	/*
		Realizamos una consulta a la entidad correspondiente con ayuda del ORM y a su vez se realizan los filtros correspondientes para obtener solo la informacion que le corresponda al usuario que realiza la peticion

		We make a query to the corresponding entity with the help of the ORM and at the same time the corresponding filters are made to obtain only the information that corresponds to the user who makes the request.
	*/
	if err := _repo_.DB.Find(&todos).Error; err != nil {
		/*
			En caso de que la consulta retorne un Error retornamos la estructura de la entidad en nula y el error resultante

			In case the query returns an Error we return the structure of the entity in null and the resulting error.
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la consulta previa y error en nulo

		If no detail is presented, the data obtained in the previous consultation will be returned and there will be no error.
	*/
	return todos, nil
}
func (_repo_ *TodoRepository) GetAll(_ctx context.Context) ([]*entity.Todo, error) {
	/*
		Se declara la siguiente variable a la cual se le asignara la estructura de la entidad

		The following variable is declared to which the structure of the entity will be assigned
	*/
	todos := []*entity.Todo{}

	/*
		Realizamos una consulta a la entidad correspondiente con ayuda del ORM y a su vez se realizan los filtros correspondientes para obtener solo la informacion que le corresponda al usuario que realiza la peticion

		We make a query to the corresponding entity with the help of the ORM and at the same time the corresponding filters are made to obtain only the information that corresponds to the user who makes the request.
	*/

	if err := _repo_.DB.Unscoped().Find(&todos).Error; err != nil {
		/*
			En caso de que la consulta retorne un Error retornamos la estructura de la entidad en nula y el error resultante

			In case the query returns an Error we return the structure of the entity in null and the resulting error.
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la consulta previa y error en nulo

		If no detail is presented, the data obtained in the previous consultation will be returned and there will be no error.
	*/

	return todos, nil

}

func (_repo_ *TodoRepository) GetByID(ctx context.Context, id uint) (*entity.Todo, error) {
	/*
		Se declara la siguiente variable a la cual se le asignara la estructura de la entidad

		The following variable is declared to which the structure of the entity will be assigned
	*/
	todo := &entity.Todo{}

	/*
		Realizamos una consulta a la entidad correspondiente con ayuda del ORM y a su vez se realizan los filtros correspondientes para obtener solo la informacion que le corresponda al usuario que realiza la peticion

		We make a query to the corresponding entity with the help of the ORM and at the same time the corresponding filters are made to obtain only the information that corresponds to the user who makes the request.
	*/
	if err := _repo_.DB.
		First(&todo, id).Error; err != nil {
		/*
			En caso de que la consulta retorne un Error retornamos la estructura de la entidad en nula y el error resultante

			In case the query returns an Error we return the structure of the entity in null and the resulting error.
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la consulta previa y error en nulo

		If no detail is presented, the data obtained in the previous consultation will be returned and there will be no error.
	*/
	return todo, nil
}

func (_repo_ *TodoRepository) Filters(ctx context.Context, _filter *utils.Filters) ([]*entity.Todo, error) {

	/*
		Se declara la siguiente variable a la cual se le asignara la estructura de la entidad

		The following variable is declared to which the structure of the entity will be assigned
	*/
	todos := []*entity.Todo{}

	/*
		Se hace una peticion al metodo NewFilter para que se realicen los filtros correspondientes con ayuda de GORM a la entidad correspondiente

		A request is made to the NewFilter method so that the corresponding filters are carried out with the help of GORM to the corresponding entity.

	*/

	if err := utils.NewFilter(&todos, _filter, _repo_.DB); err != nil {
		/*
			En caso de que la consulta retorne un Error retornamos la estructura de la entidad en nula y el error resultante

			In case the query returns an Error we return the structure of the entity in null and the resulting error.
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la consulta previa y error en nulo

		If no detail is presented, the data obtained in the previous consultation will be returned and there will be no error.
	*/
	return todos, nil
}
func (_repo_ *TodoRepository) Create(ctx context.Context, _todo *entity.Todo) (uint, error) {

	/*
		Inicializamos la variable tx la cual almacenara una transaccion de la base de datos

		We initialize the tx variable which will store a database transaction.
	*/
	tx := _repo_.DB.Begin()
	// Note the use of tx as the database handle once you are within a transaction

	/*
		Validamos que se haya asignado correctamente la transaccion

		We validate that the transaction has been assigned correctly
	*/
	if tx.Error != nil {
		return 0, tx.Error
	}

	/*
		Hacemos una peticion al metodo Create de GORM para guradar la informacion obtenida a la base de datos

		We make a request to the GORM Create method to save the information obtained to the database.
	*/
	if err := tx.Create(_todo).Error; err != nil {
		/*
			En caso de que surja un detalle en la transaccion se ejecuta la instruccion rollback para que se desagan los
			cambios en la base de datos

			In case a detail arises in the transaction, the rollback instruction is executed so that you can unpack the data.
			changes to the database
		*/
		tx.Rollback()
		return 0, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la consulta previa y error ejecutamos el
		Commit para que se realicen los cambios de la transaccion

		If no detail is presented, the data obtained in the previous consultation will be returned and the error will be executed.
		Commit for the transaction changes to be made
	*/
	return _todo.ID, tx.Commit().Error
}
func (_repo_ *TodoRepository) Update(ctx context.Context, _todo *entity.Todo) (uint, error) {

	/*
		Inicializamos la variable tx la cual almacenara una transaccion de la base de datos

		We initialize the tx variable which will store a database transaction.
	*/
	tx := _repo_.DB.Begin()
	// Note the use of tx as the database handle once you are within a transaction

	/*
		Validamos que se haya asignado correctamente la transaccion

		We validate that the transaction has been assigned correctly
	*/
	if tx.Error != nil {
		return 0, tx.Error
	}
	/*
		Hacemos una peticion al metodo Save de GORM para guradar la informacion obtenida a la base de datos

		We make a request to the GORM Save method to save the information obtained to the database.
	*/
	if err := tx.Save(&_todo).Error; err != nil {
		/*
			En caso de que surja un detalle en la transaccion se ejecuta la instruccion rollback para que se desagan los
			cambios en la base de datos

			In case a detail arises in the transaction, the rollback instruction is executed so that you can unpack the data.
			changes to the database
		*/
		tx.Rollback()
		return 0, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la consulta previa y error ejecutamos el
		Commit para que se realicen los cambios de la transaccion

		If no detail is presented, the data obtained in the previous consultation will be returned and the error will be executed.
		Commit for the transaction changes to be made
	*/
	return _todo.ID, tx.Commit().Error

}
func (_repo_ *TodoRepository) Delete(ctx context.Context, _todo *entity.Todo) (bool, error) {

	/*
		Inicializamos la variable tx la cual almacenara una transaccion de la base de datos

		We initialize the tx variable which will store a database transaction.
	*/
	tx := _repo_.DB.Begin()
	// Note the use of tx as the database handle once you are within a transaction

	/*
		Validamos que se haya asignado correctamente la transaccion

		We validate that the transaction has been assigned correctly
	*/
	if tx.Error != nil {
		return false, tx.Error
	}

	/*
		Hacemos una peticion al metodo Save de GORM para guradar la informacion obtenida a la base de datos

		We make a request to the GORM Save method to save the information obtained to the database.
	*/
	if err := tx.Save(&_todo).Error; err != nil {
		/*
			En caso de que surja un detalle en la transaccion se ejecuta la instruccion rollback para que se desagan los
			cambios en la base de datos

			In case a detail arises in the transaction, the rollback instruction is executed so that you can unpack the data.
			changes to the database
		*/
		tx.Rollback()
		return false, err
	}

	/*
		Eliminamos el registro del cual recivimos la peticion
		We deleted the record from which we received the petition.
	*/

	if err := _repo_.DB.Delete(&_todo).Error; err != nil {

		/*
		   En caso de que surja un detalle en la transaccion se ejecuta la instruccion rollback para que se desagan los
		   cambios en la base de datos

		   In case a detail arises in the transaction, the rollback instruction is executed so that you can unpack the data.
		   changes to the database
		*/
		tx.Rollback()
		/*
			Retornamos false en caso de que surgiera un detalle en la peticion

			We return false in case a detail arises in the petition.
		*/
		return false, err
	}
	/*
		Si no se presenta ningun detalle retornamos un true en caso de que el registro se elimine de manera correcta y
		error ejecutamos el Commit para que se realicen los cambios de la transaccion

		If no detail is presented we return a true in case the record is deleted correctly and
		error we execute the Commit so that the changes of the transaction are made
	*/

	return true, tx.Commit().Error
}
