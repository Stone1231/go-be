package middlewaves

import (
	"log"
	"net/http"
	"strings"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type SysMiddlewave struct {
}

var Sys = new(SysMiddlewave)

func (self *SysMiddlewave) NoRoute(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

func (self *SysMiddlewave) Recover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			// that recovery also handle XHR's
			// you need handle it
			if self.xhr(c) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": rec,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"state": -1,
					"msg":   rec,
				})
			}
		}
	}(c)
	c.Next()
}

func (self *SysMiddlewave) xhr(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}
