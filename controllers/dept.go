package controllers

import (
	Model "go-be/models"
	"go-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeptController struct {
}

var Dept = new(DeptController)

func (t *DeptController) GetAll(c *gin.Context) {
	list := services.Dept.ReadAll()
	c.JSON(http.StatusOK, list)
}

func (t *DeptController) Get(c *gin.Context) {
	var model Model.Dept

	if err := c.ShouldBindUri(&model); err != nil {
		c.JSON(400, err.Error())
		return
	}

	model = services.Dept.ReadOne(model.ID)

	if model.ID == 0 {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	c.JSON(http.StatusOK, model)
}

func (t *DeptController) Delete(c *gin.Context) {
	var model Model.Dept
	// uri/:id
	if err := c.ShouldBindUri(&model); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := services.Dept.Delete(model.ID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
