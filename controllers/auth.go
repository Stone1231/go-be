package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	mw "go-be/middlewaves"
	Model "go-be/models"
)

type AuthController struct {
}

var Auth = new(AuthController)

func (t *AuthController) Get(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(mw.Auth.IdentityKey)
	c.JSON(200, gin.H{
		mw.Auth.IdentityKey: claims[mw.Auth.IdentityKey],
		"username":          user.(*Model.AdminUser).UserName,
		"role":              user.(*Model.AdminUser).Role,
	})
}
