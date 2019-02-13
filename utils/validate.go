package utils

/*
Importamos las librerias

We import libraries
 */
import (
	validator "gopkg.in/go-playground/validator.v9"
)

/*
Funcion que valida que los datos seteados en una estructura sea correcta

Function that validates that the data set in a structure is correct
 */
func RequestValid(_struct interface{}) (bool, error) {
	/*
	Declaramos la variable validate la cual es una nueva instancia de la libreria validator

	We declare the variable validate which is a new instance of the library validator
	 */
	validate := validator.New()

	/*
	Validamos la estructura que recivimos por parametro

	Validate the structure we received by parameter
	 */
	err := validate.Struct(_struct)
	if err != nil {
		return false, err
	}
	return true, nil
}
