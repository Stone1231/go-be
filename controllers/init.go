package controllers

import (
	"go-be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InitController struct {
}

var Init = new(InitController)

func (t *InitController) All(c *gin.Context) {

	if err := services.Init.All(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *InitController) Dept(c *gin.Context) {

	if err := services.Dept.Init(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *InitController) Proj(c *gin.Context) {

	if err := services.Proj.Init(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *InitController) User(c *gin.Context) {

	if err := services.User.Init(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *InitController) Clear(c *gin.Context) {

	if err := services.Init.Clear(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *InitController) DeptClear(c *gin.Context) {

	if err := services.Dept.Clear(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *InitController) ProjClear(c *gin.Context) {

	if err := services.Proj.Clear(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *InitController) UserClear(c *gin.Context) {

	if err := services.User.Clear(); err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
