package controller

import (
	application "software/src/Application"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetByIdController struct {
	usecase application.GetByIdUseCase
}

func NewGetByIdController(app application.GetByIdUseCase)*GetByIdController{
	return &GetByIdController{
		usecase: app,
	}
}

func (c *GetByIdController) GetAreaById(ctx *gin.Context) {
    idParam := ctx.Param("id")

    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    area, err := c.usecase.ExecuteGetById(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if area == nil {
        ctx.JSON(http.StatusNotFound, gin.H{"message": "Maestría no encontrada"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
		"id":                 area.IdMaestria,
        "nombre_maestria":    area.Nombre_maestria,
        "description":        area.Description,
        "salario":            area.Salario,
        "areas":              area.Areas,
        "cantidad_escuelas":  area.Cantidad_escuelas,
        "escuelas":           area.Escuelas,
        "imagen_principal":   area.Imagen_principal,
    })
}
