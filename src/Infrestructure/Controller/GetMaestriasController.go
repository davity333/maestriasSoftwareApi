package controller

import (
	application "software/src/Application"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetAllAreasController struct {
	usecase application.GetAreasUseCase
}

func NewGetAreasController(app application.GetAreasUseCase)*GetAllAreasController{
	return &GetAllAreasController{
		usecase: app,
	}
}

func (c *GetAllAreasController)GetAreasController(ctx *gin.Context){
	areas, err := c.usecase.ExecuteGetAll()

		if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Details": "Error al obtener los episodios",
			"Error": err.Error(),
		})
		return
	}

	var dataAreas []gin.H

	for _, area := range areas{
		data := gin.H{
			"id": area.IdMaestria,
			"name": area.Nombre_maestria,
			"imagen_miniatura": area.Imagen_miniatura,
			"salario": area.Salario,
			"cantidad_escuelas": area.Cantidad_escuelas,
		}
		dataAreas = append(dataAreas, data)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Areas del area de Software": dataAreas,
	})

}