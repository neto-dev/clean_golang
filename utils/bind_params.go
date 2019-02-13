package utils

/*
Importamos las librerias

We import libraries
 */
import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
)

/*
Funcion que se utiliza para setear los valores provenientes del body de un request en la estructura que se desee

Function used to set the values coming from the body of a request in the desired structure.
 */
func BindParams(_ctx echo.Context, s interface{}) error {

	/*
	Obtenemos los valores del body

	We get the values of the body
	 */
	body, err := _ctx.Request().GetBody()
	if err != nil {
		/*
		En caso de que al obtener los valores del body surge un error lo retornamos

		In case that when obtaining the values of the body an error arises we return it
		 */
		return err
	}

	/*
	Convertimos los valores del body a bytes

	Convert body values to bytes
	 */
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		/*
		Si al momento de convertir a bytes surge un error se retorna nil en este apartado

		If at the moment of converting to bytes an error arises, it is returned nil in this section.
		 */
		return nil
	}
	if len(bytes) == 0 {
		/*
		Si el ancho del byte es igual a cero se retorna en nulo

		If the byte width is equal to zero, it is returned in null.
		 */
		return nil
	}

	/*
	Convertimos los bytes a formato json

	Convert bytes to json format
	 */
	data := json.Unmarshal(bytes, s)

	/*
	Retornamos los datos

	We return the data
	 */
	return data
}
