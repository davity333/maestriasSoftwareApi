package dependencies

import (
	application "software/src/Application"
	controller "software/src/Infrestructure/Controller"
	repositories "software/src/Domain/Repositories"
	routes "software/src/Infrestructure/Routes"

	"github.com/gin-gonic/gin"
)

func InitMaestria() *gin.Engine{
	r := gin.Default()

	maestriaRepo := repositories.NewAreaStaticRepository()

	getAaestriaUseCase := application.NewGetAreasUseCase(maestriaRepo)
	getByIdmaestriaUseCase := application.NewGetByIdUseCase(maestriaRepo)

	maestriaController := controller.NewGetAreasController(*getAaestriaUseCase)
	maestriaByIdController := controller.NewGetByIdController(*getByIdmaestriaUseCase)

	routes.MaestriaRoute(r, maestriaController, maestriaByIdController)

	return r

}