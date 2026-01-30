package routes

import (
	controller "software/src/Infrestructure/Controller"

	"github.com/gin-gonic/gin"
)

func MaestriaRoute(g *gin.Engine, getAllController *controller.GetAllAreasController,
getByIdController *controller.GetByIdController){
	router := g.Group("/api/v1/maestrias")

	router.GET("/getAll", getAllController.GetAreasController)
	router.GET("/getById/:id", getByIdController.GetAreaById)

}