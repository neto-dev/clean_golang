package utils

import "github.com/jinzhu/gorm"

type Filters struct {
	Where   string `json:"where"`
	Order   string `json:"order"`
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Select  string `json:"select"`
	Join    string `json:"join"`
	Include []struct {
		Schema string `json:"schema"`
		Where  string `json:"where"`
	} `json:"include"`
}

func NewFilter(_i interface{}, _params *Filters, _DB *gorm.DB) error {

	/*
		If within the parameters of the request the Query
		parameter is different from null it indicates that
		we will filter the results depending on the Query
		obtained.

		Si dentro de los parametros del request el parametro
		Query es diferente a null nos indica que filtraremos
		los resultados dependiendo el Query obtenido
	*/
	if _params.Where != "" {
		/*
			If within the parameters of the request the
			Order parameter is different from null it
			indicates that we will order the results
			according to the obtained request.

			Si dentro de los parámetros del request el
			parámetro Order es diferente a null nos
			indica que ordenaremos los resultados según
			la petición obtenida
		*/
		if _params.Order != "" {
			/*
				If within the parameters of the request
				the parameter Limit is different from null
				it indicates how many are the results that
				will be obtained according to the obtained
				request.

				Si dentro de los parámetros del request el
				parámetro Limit es diferente a null nos
				indica cuantos son los resultados que se
				van a obtener según la petición obtenida
			*/
			if _params.Limit != 0 {
				/*
					If within the parameters of the request
					the Page parameter is different from null,
					it indicates that pagination is being used
					and the information will be returned according
					to the page number you want to obtain.

					Si dentro de los parámetros del request el
					parámetro Page es diferente a null nos indica
					que se esta usando paginación y se devolverá
					la información según el numero de pagina que
					quiera obtener
				*/
				if _params.Page != 0 {
					if _params.Page == 1 {
						/*
							If the value of Page is equal to 1 we
							will return the values from register 0
							of the query depending on the received
							limit, at this point all the filters are
							carried out since it is detected that it
							complies with all the parameters to arrive
							at this point, the results will be returned
							according to the parameters received from
							the request.

							Si el valor de Page es igual a 1 retornaremos
							los valores a partir del registro 0 de la
							consulta dependiendo del limit recibido,
							en este punto se realizan todos los filtros
							ya que se detecto que cumple con todos los
							parámetros para llegar a este punto, se
							retornara los resultados según los parámetros
							recibidos desde el request.
						*/

						if len(_params.Include) > 0 {

							data := _DB
							if _params.Join != "" {
								data.Joins(_params.Join)
							}
							for field := 0; field < len(_params.Include); field++ {
								data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i)
							}
							return data.Error
						}
						if _params.Join != "" {
							return _DB.Joins(_params.Join).Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i).Error
						}
						return _DB.Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i).Error
					}
					/*
						If the value of Page is equal to 2 we
						will return the values from register 0
						of the query depending on the received
						limit, at this point all the filters are
						carried out since it is detected that it
						complies with all the parameters to arrive
						at this point, the results will be returned
						according to the parameters received from
						the request.

						Si el valor de Page es igual a 2 retornaremos
						los valores a partir del registro 0 de la
						consulta dependiendo del limit recibido,
						en este punto se realizan todos los filtros
						ya que se detecto que cumple con todos los
						parámetros para llegar a este punto, se
						retornara los resultados según los parámetros
						recibidos desde el request.
					*/
					if _params.Page >= 2 {
						/*
							We assign in variable the values of Page
							and Limit

							Asignamos en variable los valores de
							Page y Limit
						*/
						page := _params.Page
						limit := _params.Limit
						/*
							We obtain the offset based on the page
							value and limit the way in which the
							correct registers are obtained is by
							subtracting 1 from the value received
							from Page _params and then multiplying
							it by the value of Limit _params.
							Example:
							page: = 2
							limit: =10
							offset = (page - 1) * (limit) This returns 1

							Obtenemos el offset en base a el valor
							de page y de limit la manera en como
							se obtienen los registros correctos es
							restando 1 al valor que se recibe de
							_params Page y posteriormente
							multiplicándolo por el valor de _params
							Limit.
							Ejemplo:
							page := 2
							limit :=10
							offset = (page - 1) * (limit) Esto retorna 10 0
						*/
						offset := (page - 1) * (limit)
						/*
							We return the result

							Retornamos el resultado
						*/

						if len(_params.Include) > 0 {

							data := _DB
							if _params.Join != "" {
								data.Joins(_params.Join)
							}
							for field := 0; field < len(_params.Include); field++ {
								data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i)
							}
							return data.Error
						}
						if _params.Join != "" {
							return _DB.Joins(_params.Join).Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
						}
						return _DB.Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
					}
				}
				/*
					In case Page comes empty we return the result
					without applying that filter.

					En caso de que Page venga vacío retornamos el
					resultado sin aplicar ese filtro
				*/

				if len(_params.Include) > 0 {

					data := _DB
					if _params.Join != "" {
						data.Joins(_params.Join)
					}
					for field := 0; field < len(_params.Include); field++ {
						data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i)
					}
					return data.Error
				}
				if _params.Join != "" {
					return _DB.Joins(_params.Join).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
				}
				return _DB.Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
			}

			if len(_params.Include) > 0 {

				data := _DB
				if _params.Join != "" {
					data.Joins(_params.Join)
				}
				for field := 0; field < len(_params.Include); field++ {
					data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Order(_params.Order).Where(_params.Where).Find(_i)
				}
				return data.Error
			}
			if _params.Join != "" {
				return _DB.Joins(_params.Join).Order(_params.Order).Where(_params.Where).Find(_i).Error
			}
			return _DB.Order(_params.Order).Where(_params.Where).Find(_i).Error
		}

		/*
			Other conditions interact in the same way validate
			the data being received and return a response as
			the case may be.

			El Resto de condiciones interactúan de la misma
			forma validan los datos que se están recibiendo
			y retorna una respuesta según sea el caso.
		*/

		if _params.Limit != 0 {
			if _params.Page != 0 {
				if _params.Page == 1 {
					if len(_params.Include) > 0 {

						data := _DB
						if _params.Join != "" {
							data.Joins(_params.Join)
						}
						for field := 0; field < len(_params.Include); field++ {
							data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i)
						}
						return data.Error
					}
					if _params.Join != "" {
						return _DB.Joins(_params.Join).Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i).Error
					}
					return _DB.Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i).Error
				}
				if _params.Page >= 2 {
					page := _params.Page
					limit := _params.Limit
					offset := (page - 1) * (limit)
					if len(_params.Include) > 0 {

						data := _DB
						if _params.Join != "" {
							data.Joins(_params.Join)
						}
						for field := 0; field < len(_params.Include); field++ {
							data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i)
						}
						return data.Error
					}
					if _params.Join != "" {
						return _DB.Joins(_params.Join).Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
					}
					return _DB.Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
				}
			}
			if len(_params.Include) > 0 {

				data := _DB
				if _params.Join != "" {
					data.Joins(_params.Join)
				}
				for field := 0; field < len(_params.Include); field++ {
					data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Limit(_params.Limit).Where(_params.Where).Find(_i)
				}
				return data.Error
			}
			if _params.Join != "" {
				return _DB.Joins(_params.Join).Limit(_params.Limit).Where(_params.Where).Find(_i).Error
			}
			return _DB.Limit(_params.Limit).Where(_params.Where).Find(_i).Error
		}

		if len(_params.Include) > 0 {

			data := _DB
			if _params.Join != "" {
				data.Joins(_params.Join)
			}
			for field := 0; field < len(_params.Include); field++ {
				data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Where(_params.Where).Find(_i)
			}
			return data.Error
		}

		if _params.Join != "" {
			return _DB.Joins(_params.Join).Where(_params.Where).Find(_i).Error
		}

		return _DB.Where(_params.Where).Find(_i).Error
	}

	/*
		The page parameter is only taken into account when
		the limit parameter is also received.

		El parametro page solo se toma en cuenta cuando
		también se recibe el parámetro de limit
	*/

	if _params.Order != "" {
		if _params.Limit != 0 {
			if _params.Page != 0 {
				if _params.Page == 1 {
					if len(_params.Include) > 0 {

						data := _DB
						if _params.Join != "" {
							data.Joins(_params.Join)
						}
						for field := 0; field < len(_params.Include); field++ {
							data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i)
						}
						return data.Error
					}
					if _params.Join != "" {
						return _DB.Joins(_params.Join).Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i).Error
					}
					return _DB.Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i).Error
				}
				if _params.Page >= 2 {
					page := _params.Page
					limit := _params.Limit
					offset := (page - 1) * (limit)
					if len(_params.Include) > 0 {

						data := _DB
						if _params.Join != "" {
							data.Joins(_params.Join)
						}
						for field := 0; field < len(_params.Include); field++ {
							data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i)
						}
						return data.Error
					}
					if _params.Join != "" {
						return _DB.Joins(_params.Join).Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
					}
					return _DB.Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
				}
			}
			if len(_params.Include) > 0 {

				data := _DB
				if _params.Join != "" {
					data.Joins(_params.Join)
				}
				for field := 0; field < len(_params.Include); field++ {
					data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Limit(_params.Limit).Order(_params.Order).Find(_i)
				}
				return data.Error
			}
			if _params.Join != "" {
				return _DB.Joins(_params.Join).Limit(_params.Limit).Order(_params.Order).Find(_i).Error
			}
			return _DB.Limit(_params.Limit).Order(_params.Order).Find(_i).Error
		}
		if len(_params.Include) > 0 {

			data := _DB
			if _params.Join != "" {
				data.Joins(_params.Join)
			}
			for field := 0; field < len(_params.Include); field++ {
				data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Order(_params.Order).Find(_i)
			}
			return data.Error
		}
		if _params.Join != "" {
			return _DB.Joins(_params.Join).Order(_params.Order).Find(_i).Error
		}
		return _DB.Order(_params.Order).Find(_i).Error
	}

	if _params.Limit != 0 {
		if _params.Page != 0 {
			if _params.Page == 1 {
				if len(_params.Include) > 0 {

					data := _DB
					if _params.Join != "" {
						data.Joins(_params.Join)
					}
					for field := 0; field < len(_params.Include); field++ {
						data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i)
					}
					return data.Error
				}
				if _params.Join != "" {
					return _DB.Joins(_params.Join).Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i).Error
				}
				return _DB.Limit(_params.Limit).Offset(0).Order(_params.Order).Where(_params.Where).Find(_i).Error
			}
			if _params.Page >= 2 {
				page := _params.Page
				limit := _params.Limit
				offset := (page - 1) * (limit)
				if len(_params.Include) > 0 {

					data := _DB
					if _params.Join != "" {
						data.Joins(_params.Join)
					}
					for field := 0; field < len(_params.Include); field++ {
						data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i)
					}
					return data.Error
				}
				if _params.Join != "" {
					return _DB.Joins(_params.Join).Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
				}
				return _DB.Offset(offset).Limit(_params.Limit).Order(_params.Order).Where(_params.Where).Find(_i).Error
			}
		}
		if len(_params.Include) > 0 {

			data := _DB
			for field := 0; field < len(_params.Include); field++ {
				data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Limit(_params.Limit).Find(_i)
			}
			return data.Error
		}
		return _DB.Limit(_params.Limit).Find(_i).Error
	}

	if len(_params.Include) > 0 {

		data := _DB
		for field := 0; field < len(_params.Include); field++ {
			data = data.Preload(_params.Include[field].Schema, _params.Include[field].Where).Find(_i)
		}
		return data.Error
	}

	if _params.Select != "" {
		return _DB.Select(_params.Select).Find(_i).Error
	}

	if _params.Join != "" {
		return _DB.Joins(_params.Join).Find(_i).Error
	}

	/*
		If no data is received some returns all records.

		Si no se reciben datos algunos retorna todos los registros.
	*/
	return _DB.Find(_i).Error
}
