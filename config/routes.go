/*
 * Project: dhq_server
 * Developer: Ernesto Valenzuela Vargas.
 * Created by: netodev
 * Contact: neto.dev@protonmail.com
 *
 */

/*
 * Revision History:
 *     Initial:      10/18/18  |  1:56 PM     Ernesto Valenzuela V
 */

package config

/*
Importamos las librerias

We import libraries
*/
import (
	"github.com/clean_golang/data/repository_impl"
	"github.com/clean_golang/domain/usecase_impl"
	"github.com/clean_golang/presentation/http"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

/*
Metodo encargado de crear las rutas de cada Handler y de inyectar las dependencias de cada capa perteneciente a cada
entidad

Method in charge of creating the routes of each Handler and injecting the dependencies of each layer belonging to each Handler.
entity
*/
func Routes(DB *gorm.DB, group echo.Group) {
	// >>>>>>>>>>>> todoS
	/*
		Declaramos el repositorio de la entidad y le pasamos la base de datos correspondiente

		Declare the repository of the entity and pass it the corresponding database
	*/
	todoRepository := repository_impl.NewTodoRepository(DB)
	/*
		Declaramos el caso de uso de la entidad actual y le inyectamos el repositorio que se desea implementar

		Declare the case of use of the current entity and inject the repository you want to implement
	*/
	todoUsecase := usecase_impl.NewTodoUsecase(todoRepository)
	/*
		Pasamos al heandler de la entidad el grupo de las rutas y le inyectamos la dependencia del caso de uso

		We pass to the heandler of the entity the group of the routes and we inject to him the dependence of the case of use
	*/
	http.NewTodoHandler(group, todoUsecase)

}
