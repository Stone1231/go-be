package controllers

import (
	"fmt"
	Model "go-be/models"
	"go-be/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var User = new(UserController)

// @Tags user
// @Produce json
// @Param user body models.User true "create user"
// @Success 200 {object} Model.User
// @Router /api/user [post]
func (t *UserController) Post(c *gin.Context) {
	var model Model.User

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	model.Birthday, _ = time.Parse("2006-01-02", model.BirthdayStr)

	if err := services.User.Create(&model); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model)
}

func (t *UserController) Put(c *gin.Context) {
	var model Model.User

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// uri/:id
	if err := c.ShouldBindUri(&model); err != nil {
		c.JSON(400, err.Error())
		return
	}

	model.Birthday, _ = time.Parse("2006-01-02", model.BirthdayStr)

	if err := services.User.Update(&model); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model)
}

// @Tags user
// @Produce json
// @Success 200 {object} models.User
// @Router /api/user/{id} [get]
func (t *UserController) Get(c *gin.Context) {
	var model Model.User

	if err := c.ShouldBindUri(&model); err != nil {
		c.JSON(400, err.Error())
		return
	}

	model = services.User.ReadOne(model.ID)

	c.JSON(http.StatusOK, model)
}

func (t *UserController) GetAll(c *gin.Context) {
	list := services.User.ReadAll()
	c.JSON(http.StatusOK, list)
}

func (t *UserController) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename, err := services.File.Upload(&file, header.Filename, "img")

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, filename)
}

func (t *UserController) Query(c *gin.Context) {
	var name string
	c.Bind(&name)
	list := services.User.Query(name)
	c.JSON(http.StatusOK, list)
}

func (t *UserController) Delete(c *gin.Context) {
	var model Model.User
	// uri/:id
	if err := c.ShouldBindUri(&model); err != nil {
		c.JSON(400, err.Error())
		return
	}

	if err := services.User.Delete(model.ID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
