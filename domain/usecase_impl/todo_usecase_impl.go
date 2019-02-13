package usecase_impl

/*
Importamos las librerias

We import libraries
*/
import (
	"context"

	"github.com/clean_golang/data/entity"
	"github.com/clean_golang/data/repository"
	"github.com/clean_golang/domain/usecase"
	"github.com/clean_golang/utils"
)

/*
Estructura la cual almacena el repositorio con el que va a trabajar en el momento de ejecucion,
todas las funciones dependeran de esta para poder acceder a los metodos del repositorio.

Structure which stores the repository with which it will work at the time of execution,
all functions will depend on it to access the repository methods.
*/
type todoUsecase struct {
	todoRepos repository.TodoRepository
}

/*
Funcion encargada de crear una nueva instancia del caso de uso para poder realizar la inyeccion de dependencias

Function in charge of creating a new instance of the use case to be able to carry out the injection of dependencies.
*/

func NewTodoUsecase(a repository.TodoRepository) usecase.TodoUsecase {
	/*
		Retornamos la instancia del caso de uso junto con la inyeccion del repositorio correspondiente

		We return the instance of the use case together with the injection of the corresponding repository.
	*/
	return &todoUsecase{
		todoRepos: a,
	}
}

/*
Funcion encargada de consultar el metodo Get de la capa de repositorio

Function in charge of consulting the Get method of the repository layer
*/
func (_uc_ *todoUsecase) Get(_ctx context.Context) ([]*entity.Todo, error) {
	/*
		Se declaran dos variables res y err la primera almacenara la respuesta del repositorio,
		y err almacena el error en caso de que surgiera alguno

		Two variables are declared res and err the first one will store the repository response,
		and err stores the error in case one arises
	*/
	res, err := _uc_.todoRepos.Get(_ctx)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la peticion previa y error en nulo

		If no details are presented, the data obtained in the previous request will be returned and the error will be null and void.
	*/

	return res, nil
}

/*
Funcion encargada de consultar el metodo Getall de la capa de repositorio

Function in charge of consulting the GetAll method of the repository layer
*/
func (_uc_ *todoUsecase) GetAll(_ctx context.Context) ([]*entity.Todo, error) {

	/*
		Se declaran dos variables res y err la primera almacenara la respuesta del repositorio,
		y err almacena el error en caso de que surgiera alguno

		Two variables are declared res and err the first one will store the repository response,
		and err stores the error in case one arises
	*/
	res, err := _uc_.todoRepos.GetAll(_ctx)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la peticion previa y error en nulo

		If no details are presented, the data obtained in the previous request will be returned and the error will be null and void.
	*/

	return res, nil
}

/*
Funcion encargada de consultar el metodo Filters de la capa de repositorio

Function in charge of consulting the Filters method of the repository layer
*/
func (_uc_ *todoUsecase) Filters(_ctx context.Context, _filter *utils.Filters) ([]*entity.Todo, error) {
	/*
		Se declaran dos variables res y err la primera almacenara la respuesta del repositorio,
		y err almacena el error en caso de que surgiera alguno

		Two variables are declared res and err the first one will store the repository response,
		and err stores the error in case one arises
	*/
	res, err := _uc_.todoRepos.Filters(_ctx, _filter)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la peticion previa y error en nulo

		If no details are presented, the data obtained in the previous request will be returned and the error will be null and void.
	*/

	return res, nil
}

/*
Funcion encargada de consultar el metodo Filters de la capa de repositorio

Function in charge of consulting the Filters method of the repository layer
*/
func (_uc_ *todoUsecase) GetByID(_ctx context.Context, id uint) (*entity.Todo, error) {

	/*
		Se declaran dos variables res y err la primera almacenara la respuesta del repositorio,
		y err almacena el error en caso de que surgiera alguno

		Two variables are declared res and err the first one will store the repository response,
		and err stores the error in case one arises
	*/
	res, err := _uc_.todoRepos.GetByID(_ctx, id)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la peticion previa y error en nulo

		If no details are presented, the data obtained in the previous request will be returned and the error will be null and void.
	*/

	return res, nil
}

/*
Funcion encargada de consultar el metodo Create de la capa de repositorio

Function in charge of consulting the Create method of the repository layer
*/
func (_uc_ *todoUsecase) Create(_ctx context.Context, _todo *entity.Todo) (*entity.Todo, error) {

	/*
		Se declaran dos variables res y err la primera almacenara la respuesta del repositorio en este caso sera el id
		del registro creado el cual se utilizara en seguida para realizar otra peticion al repositorio,
		y err almacena el error en caso de que surgiera alguno

		Two variables are declared res and err the first will store the response of the repository in this case will be the id
		of the created record which will then be used to make another request to the repository,
		and err stores the error in case one arises
	*/
	res, err := _uc_.todoRepos.Create(_ctx, _todo)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return nil, err
	}

	/*
		Realizamos otra consulta al repositorio para obtener el nuevo registo creado anteriormente y lo almacenamos en la
		variable todo en caso de error seteamos el error en la variable err

		Make another query to the repository to get the new record previously created and store it in the variable todo in case of error we set the error in the variable err
	*/
	todo, err := _uc_.todoRepos.GetByID(_ctx, res)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la peticion previa y error en nulo

		If no details are presented, the data obtained in the previous request will be returned and the error will be null and void.
	*/

	return todo, nil
}

/*
Funcion encargada de consultar el metodo Update de la capa de repositorio

Function in charge of consulting the Update method of the repository layer
*/
func (_uc_ *todoUsecase) Update(_ctx context.Context, _todo *entity.Todo) (*entity.Todo, error) {
	/*
		Se declaran dos variables res y err la primera almacenara la respuesta del repositorio en este caso sera el id
		del registro editado el cual se utilizara en seguida para realizar otra peticion al repositorio,
		y err almacena el error en caso de que surgiera alguno

		Two variables are declared res and err the first will store the response of the repository in this case will be the id
		of the edited record which will then be used to make another request to the repository,
		and err stores the error in case one arises
	*/
	res, err := _uc_.todoRepos.Update(_ctx, _todo)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return nil, err
	}

	/*
		Realizamos otra consulta al repositorio para obtener el registo editado anteriormente y lo almacenamos en la
		variable todo en caso de error seteamos el error en la variable err

		Make another query to the repository to get the record previously updated and store it in the variable todo in
		case of	error we set the error in the variable err
	*/

	todo, err := _uc_.todoRepos.GetByID(_ctx, res)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return nil, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la peticion previa y error en nulo

		If no details are presented, the data obtained in the previous request will be returned and the error will be null and void.
	*/

	return todo, nil
}

/*
Funcion encargada de consultar el metodo Delete de la capa de repositorio

Function in charge of consulting the Delete method of the repository layer
*/
func (_uc_ *todoUsecase) Delete(_ctx context.Context, id uint) (bool, error) {
	/*
		Realizamos peticion al metodo GetByID para obtener el registro que se desea eliminar

		We make a request to the GetByID method to obtain the record you want to delete.
	*/
	todo, err := _uc_.todoRepos.GetByID(_ctx, id)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return false, err
	}

	/*
		Se declaran dos variables res y err la primera almacenara la respuesta del repositorio en este caso sera un boleano
		el cual nos comunica si el regustro fue eliminado correctamente o no, err almacena el error en caso de que surgiera alguno

		Two variables are declared res and err the first will store the repository response in this case will be a boleano
		which tells us whether the regustro was removed correctly or not, err stores the error in case it arose any
	*/

	res, err := _uc_.todoRepos.Delete(_ctx, todo)

	if err != nil {
		/*
			En caso de que el repositorio retorne un Error retornamos la estructura de la entidad en nula y el error
			resultante

			In case the repository returns an Error we return the structure of the entity in null and the error
			resulting
		*/
		return false, err
	}

	/*
		Si no se presenta ningun detalle se retornan los datos obtenidos en la peticion previa y error en nulo

		If no details are presented, the data obtained in the previous request will be returned and the error will be null and void.
	*/

	return res, nil
}
