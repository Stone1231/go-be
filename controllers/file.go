package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController struct {
}

var File = new(FileController)

func (t *FileController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.JSON(http.StatusOK, file.Filename)
}

func (t *FileController) Upload2(c *gin.Context) {
	file1, _ := c.FormFile("file1")
	file2, _ := c.FormFile("file2")
	c.JSON(http.StatusOK, []string{file1.Filename, file2.Filename})
}

func (t *FileController) Uploads(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]

	ss := make([]string, 0)

	for _, file := range files {
		ss = append(ss, file.Filename)
	}
	c.JSON(http.StatusOK, ss)
}
