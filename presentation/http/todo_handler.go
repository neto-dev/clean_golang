package http

/*
Importamos las librerias

We import libraries
*/
import (
	"context"
	"net/http"
	"strconv"

	"github.com/clean_golang/data/entity"
	"github.com/clean_golang/domain/usecase"
	"github.com/clean_golang/utils"
	"github.com/labstack/echo"
)

/*
Estructura la cual almacena el caso de uso con la que se va a trabajar en el momento de ejecucion,
todas las funciones dependeran de esta para poder acceder a las propiedades del caso de uso

Structure which stores the case of use with which it will work at the time of execution, all the functions will depend on this to be able to access the properties of the use case
*/
type HttpTodoHandler struct {
	TodoUsecase usecase.TodoUsecase
}

/*
Funcion encargada de interactuar con la peticion del cliente y encargada tambien de hacer una consulta tipo Get a la
base de datos, esta funcion retorna todos los registros pertenecientes a la entidad

Function responsible for interacting with the client's request and also responsible for making a Get type consultation to the
database, this function returns all records belonging to the entity
*/
func (_handler_ *HttpTodoHandler) Get(_c echo.Context) error {
	/*
		Se obtiene y setea el context de la peticion en la variable ctx

		It obtains and sets the context of the request in the variable ctx
	*/
	ctx := _c.Request().Context()
	if ctx == nil {
		/*
			En caso de que el contexto sea nulo se crea un contexto de tipo background

			If the context is null, a background context is created.
		*/
		ctx = context.Background()
	}

	/*
		Consultamos el caso de uso del metodo Get para obtener los valores de la entidad,
		en caso de retornar algun error lo almacenamos en la variable err

		We consulted the case of use of the Get method to obtain the values of the entity,
		in case of returning some error we store it in the variable err
	*/
	todos, err := _handler_.TodoUsecase.Get(ctx)

	if err != nil {
		/*
			Si la peticion retorna un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If the request returns an error we return a JSON response with the error received and an ertructure.
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Find Failure", err)
	}

	/*
		Si la peticion es correcta retornamos los datos de la entridad en formato JSON con los recibidos y con una
		ertructura establecida en el metodo ReturnRespondJSON

		If the request is correct we return the data of the entry in JSON format with the received ones and with a
		the structure established in the ReturnRespondJSON method
	*/
	return utils.ReturnRespondJSON(_c, http.StatusOK, todos)
}
func (_handler_ *HttpTodoHandler) GetAll(_c echo.Context) error {
	/*
		Se obtiene y setea el context de la peticion en la variable ctx

		It obtains and sets the context of the request in the variable ctx
	*/
	ctx := _c.Request().Context()
	if ctx == nil {
		/*
			En caso de que el contexto sea nulo se crea un contexto de tipo background

			If the context is null, a background context is created.
		*/
		ctx = context.Background()
	}

	/*
		Consultamos el caso de uso del metodo GetAll para obtener los valores de la entidad,
		en caso de retornar algun error lo almacenamos en la variable err

		We consulted the case of use of the GetAll method to obtain the values of the entity,
		in case of returning some error we store it in the variable err
	*/
	todos, err := _handler_.TodoUsecase.GetAll(ctx)

	if err != nil {
		/*
			Si la peticion retorna un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If the request returns an error we return a JSON response with the error received and an ertructure.
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Find Failure", err)
	}

	/*
		Si la peticion es correcta retornamos los datos de la entridad en formato JSON con los recibidos y con una
		ertructura establecida en el metodo ReturnRespondJSON

		If the request is correct we return the data of the entry in JSON format with the received ones and with a
		the structure established in the ReturnRespondJSON method
	*/
	return utils.ReturnRespondJSON(_c, http.StatusOK, todos)
}

func (_handler_ *HttpTodoHandler) Filters(_c echo.Context) error {

	/*
		Declaramos la variable filters del tipo Filters

		We declare the variable filters of the type Filters
	*/
	var filters *utils.Filters

	/*
		Seteamos los datos que recibimos en el body de la peticion en el objeto filters

		We seteamos the data that we receive in the body of the petition in the object filters
	*/
	err := _c.Bind(&filters)
	if err != nil {
		/*
			Si al setear la informacion da un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If when setting the information gives an error we return a JSON response with the error received and with an ertructura
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Unprocessable Entity", err)
	}

	/*
		Se obtiene y setea el context de la peticion en la variable ctx

		It obtains and sets the context of the request in the variable ctx
	*/
	ctx := _c.Request().Context()
	if ctx == nil {
		/*
			En caso de que el contexto sea nulo se crea un contexto de tipo background

			If the context is null, a background context is created.
		*/
		ctx = context.Background()
	}

	/*
		Consultamos el caso de uso del metodo Filters para obtener los valores de la entidad,
		en caso de retornar algun error lo almacenamos en la variable err

		We consulted the case of use of the Filters method to obtain the values of the entity,
		in case of returning some error we store it in the variable err
	*/
	todos, err := _handler_.TodoUsecase.Filters(ctx, filters)

	if err != nil {
		/*
			Si la peticion retorna un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If the request returns an error we return a JSON response with the error received and an ertructure.
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Find Failure", err)
	}
	/*
		Si la peticion es correcta retornamos los datos de la entridad en formato JSON con los recibidos y con una
		ertructura establecida en el metodo ReturnRespondJSON

		If the request is correct we return the data of the entry in JSON format with the received ones and with a
		the structure established in the ReturnRespondJSON method
	*/
	return utils.ReturnRespondJSON(_c, http.StatusOK, todos)
}

func (_handler_ *HttpTodoHandler) GetByID(_c echo.Context) error {

	/*
		Obtenemos y transformamos el parametro id que recibimos en la peticion

		We obtain and transform the id parameter that we receive in the petition.
	*/
	idP, err := strconv.Atoi(_c.Param("id"))
	/*
		Seteamos la variable idP en la variable id transformandola en uint

		We set the variable idP in the variable id transforming it into uint
	*/
	id := uint(idP)

	/*
		Se obtiene y setea el context de la peticion en la variable ctx

		It obtains and sets the context of the request in the variable ctx
	*/
	ctx := _c.Request().Context()
	if ctx == nil {
		/*
			En caso de que el contexto sea nulo se crea un contexto de tipo background

			If the context is null, a background context is created.
		*/
		ctx = context.Background()
	}

	/*
		Consultamos el caso de uso del metodo GetByID para obtener los valores de la entidad,
		en caso de retornar algun error lo almacenamos en la variable err

		We consulted the case of use of the GetByID method to obtain the values of the entity,
		in case of returning some error we store it in the variable err
	*/

	todo, err := _handler_.TodoUsecase.GetByID(ctx, id)

	if err != nil {
		/*
			Si la peticion retorna un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If the request returns an error we return a JSON response with the error received and an ertructure.
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Find Failure", err)
	}
	/*
		Si la peticion es correcta retornamos los datos de la entridad en formato JSON con los recibidos y con una
		ertructura establecida en el metodo ReturnRespondJSON

		If the request is correct we return the data of the entry in JSON format with the received ones and with a
		the structure established in the ReturnRespondJSON method
	*/
	return utils.ReturnRespondJSON(_c, http.StatusOK, todo)
}

func (_handler_ *HttpTodoHandler) Create(_c echo.Context) error {
	/*
		Declaramos la variable todo que es de tipo entidad de Todo

		We declare the variable todo that is of type entity of Todo
	*/
	var todo entity.Todo

	/*
		Seteamos los valores que recibimos en el body del request

		We set the values that we receive in the body of the request
	*/
	err := _c.Bind(&todo)
	if err != nil {
		/*
			Si al setear la informacion da un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If when setting the information gives an error we return a JSON response with the error received and with an ertructura
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Unprocessable Entity", err)
	}

	/*
		Validamos que los datos recibidos sean correctos dependiendo la estructura

		We validate that the data received are correct depending on the structure
	*/
	if ok, err := utils.RequestValid(todo); !ok {
		/*
			Si nos retorna error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If it returns error we return a JSON response with the error received and with an ertructura
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Bad Request", err)
	}

	/*
		Se obtiene y setea el context de la peticion en la variable ctx

		It obtains and sets the context of the request in the variable ctx
	*/
	ctx := _c.Request().Context()
	if ctx == nil {
		/*
			En caso de que el contexto sea nulo se crea un contexto de tipo background

			If the context is null, a background context is created.
		*/
		ctx = context.Background()
	}

	/*
		Consultamos el caso de uso del metodo Create para crear el nuevo registro,
		en caso de retornar algun error lo almacenamos en la variable err

		Consultamos el caso de uso del metodo Create para crear el nuevo registro,
		en caso de retornar algun error lo almacenamos en la variable err
	*/
	response, err := _handler_.TodoUsecase.Create(ctx, &todo)

	if err != nil {
		/*
			Si la peticion retorna un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If the request returns an error we return a JSON response with the error received and an ertructure.
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Create Failure", err)
	}
	/*
		Si la peticion es correcta retornamos los datos de la entridad en formato JSON con los recibidos y con una
		ertructura establecida en el metodo ReturnRespondJSON

		If the request is correct we return the data of the entry in JSON format with the received ones and with a
		the structure established in the ReturnRespondJSON method
	*/
	return utils.ReturnRespondJSON(_c, http.StatusOK, response)
}

func (_handler_ *HttpTodoHandler) Update(_c echo.Context) error {
	/*
		Obtenemos y transformamos el parametro id que recibimos en la peticion

		We obtain and transform the id parameter that we receive in the petition.
	*/
	idP, err := strconv.Atoi(_c.Param("id"))
	/*
		Seteamos la variable idP en la variable id transformandola en uint

		We set the variable idP in the variable id transforming it into uint
	*/
	id := uint(idP)

	/*
		Se obtiene y setea el context de la peticion en la variable ctx

		It obtains and sets the context of the request in the variable ctx
	*/
	ctx := _c.Request().Context()
	if ctx == nil {
		/*
			En caso de que el contexto sea nulo se crea un contexto de tipo background

			If the context is null, a background context is created.
		*/
		ctx = context.Background()
	}

	/*
		Primero se obtiene el registro que la peticion ha indicado que desea editar

		First you get the record that the request has indicated you want to edit
	*/
	todo, err := _handler_.TodoUsecase.GetByID(ctx, id)

	if err != nil {
		/*
			Si la peticion retorna un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If the request returns an error we return a JSON response with the error received and an ertructure.
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Find Failure", err)
	}

	/*
		Seteamos los valores que recibimos en el body del request

		We set the values that we receive in the body of the request
	*/
	if err := _c.Bind(&todo); err != nil {
		/*
			Si al setear la informacion da un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If when setting the information gives an error we return a JSON response with the error received and with an ertructura
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Unprocessable Entity", err)
	}

	/*
		Validamos que los datos recibidos sean correctos dependiendo la estructura

		We validate that the data received are correct depending on the structure
	*/
	if ok, err := utils.RequestValid(todo); !ok {
		/*
			Si nos retorna error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If it returns error we return a JSON response with the error received and with an ertructura
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Bad Request", err)
	}

	/*
		Consultamos el caso de uso del metodo Update para actualizar los valores de la entidad,
		en caso de retornar algun error lo almacenamos en la variable err

		We consulted the case of use of the Update method to update the values of the entity,
		in case of returning some error we store it in the variable err
	*/
	response, err := _handler_.TodoUsecase.Update(ctx, todo)

	if err != nil {
		/*
			Si la peticion retorna un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If the request returns an error we return a JSON response with the error received and an ertructure.
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Update Failure", err)
	}
	/*
		Si la peticion es correcta retornamos los datos de la entridad en formato JSON con los recibidos y con una
		ertructura establecida en el metodo ReturnRespondJSON

		If the request is correct we return the data of the entry in JSON format with the received ones and with a
		the structure established in the ReturnRespondJSON method
	*/
	return utils.ReturnRespondJSON(_c, http.StatusOK, response)
}

func (_handler_ *HttpTodoHandler) Delete(_c echo.Context) error {
	/*
		Obtenemos y transformamos el parametro id que recibimos en la peticion

		We obtain and transform the id parameter that we receive in the petition.
	*/
	idP, err := strconv.Atoi(_c.Param("id"))
	/*
		Seteamos la variable idP en la variable id transformandola en uint

		We set the variable idP in the variable id transforming it into uint
	*/
	id := uint(idP)

	/*
		Se obtiene y setea el context de la peticion en la variable ctx

		It obtains and sets the context of the request in the variable ctx
	*/
	ctx := _c.Request().Context()
	if ctx == nil {
		/*
			En caso de que el contexto sea nulo se crea un contexto de tipo background

			If the context is null, a background context is created.
		*/
		ctx = context.Background()
	}
	/*
		Consultamos el caso de uso del metodo Delete para eliminar el registro,
		en caso de retornar algun error lo almacenamos en la variable err

		We consulted the case of use of the Delete method to delete the record,
		in case of returning some error we store it in the variable err
	*/
	_, err = _handler_.TodoUsecase.Delete(ctx, id)

	if err != nil {
		/*
			Si la peticion retorna un error retornamos una respuesta JSON con el error recibido y con una ertructura
			establecida en el metodo ReturnErrorJSON

			If the request returns an error we return a JSON response with the error received and an ertructure.
			established in the ReturnErrorJSON method
		*/
		return utils.ReturnErrorJSON(_c, "Delete Failure", err)
	}
	/*
		Si la peticion es correcta retornamos los datos de la entridad en formato JSON con los recibidos y con una
		ertructura establecida en el metodo ReturnRespondJSON

		If the request is correct we return the data of the entry in JSON format with the received ones and with a
		the structure established in the ReturnRespondJSON method
	*/
	return utils.ReturnRespondJSON(_c, http.StatusOK, "Delete Success")
}

func NewTodoHandler(e echo.Group, us usecase.TodoUsecase) echo.Group {
	/*
		Se declara la variable handler la cual contendra toda la estructura de nuestros heandler methods y a su vez se
		asigna el caso de uso que se utilizara en nuestra estructrura

		The handler variable is declared which will contain all the structure of our heandler methods and in turn will be
		assigns the case of use that will be used in our structure
	*/
	handler := &HttpTodoHandler{
		TodoUsecase: us,
	}

	/*
		Se declaran las rutas correspondientes a cada metodo

		The routes corresponding to each method are declared.
	*/

	e.GET("/todos", handler.Get)
	e.GET("/todos/all", handler.GetAll)
	e.GET("/todos/:id", handler.GetByID)
	e.POST("/todos/filters", handler.Filters)
	e.POST("/todos", handler.Create)
	e.PUT("/todos/:id", handler.Update)
	e.PATCH("/todos/:id", handler.Update)
	e.DELETE("/todos/:id", handler.Delete)

	return e

}
