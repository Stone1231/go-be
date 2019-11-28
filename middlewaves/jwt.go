package middlewaves

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	Model "go-be/models"
)

var identityKey = "id"

var Auth, authRrr = jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "test zone",
	Key:         []byte("secret key"),
	Timeout:     time.Hour,
	MaxRefresh:  time.Hour,
	IdentityKey: identityKey,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*Model.AdminUser); ok {
			return jwt.MapClaims{
				identityKey: v.UserName,
				"username":  v.UserName,
				"role":      v.Role,
			}
		}
		return jwt.MapClaims{}
	},
	IdentityHandler: func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &Model.AdminUser{
			UserName: claims[identityKey].(string),
			Role:     claims["role"].(string),
		}
	},
	Authenticator: func(c *gin.Context) (interface{}, error) {
		var loginVals Model.Login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		switch loginVals.Username + "-" + loginVals.Password {
		case "user-pwd":
			return &Model.AdminUser{
				UserName: loginVals.Username,
				Role:     "dev",
			}, nil
		case "test-pwd":
			return &Model.AdminUser{
				UserName: loginVals.Username,
				Role:     "test",
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	},
	Authorizator: func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*Model.AdminUser); ok && v.Role == "dev" {
			return true
		}

		return false
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	// - "param:<name>"
	TokenLookup: "header: Authorization, query: token, cookie: jwt",
	// TokenLookup: "query:token",
	// TokenLookup: "cookie:token",

	// TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName: "Bearer",

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc: time.Now,
})

// if err != nil {
// 	log.Fatal("JWT Error:" + err.Error())
// }
