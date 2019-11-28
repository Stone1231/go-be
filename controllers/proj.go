package controllers

import (
	Model "go-be/models"
	"go-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjController struct {
}

var Proj = new(ProjController)

func (t *ProjController) GetAll(c *gin.Context) {
	list := services.Proj.ReadAll()
	c.JSON(http.StatusOK, list)
}

func (t *ProjController) Delete(c *gin.Context) {
	var model Model.Proj
	// uri/:id
	if err := c.ShouldBindUri(&model); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := services.Proj.Delete(model.ID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
