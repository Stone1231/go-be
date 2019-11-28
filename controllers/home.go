package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

var Home = new(HomeController)

func (t *HomeController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, "GO webapi")
}

func (t *HomeController) Error(c *gin.Context) {
	panic("error test!")
	c.JSON(http.StatusOK, "Error Test!")
}
