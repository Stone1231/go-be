package main

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"go-be/controllers"
	"go-be/database"
	mw "go-be/middlewaves"
	"go-be/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	log.Println("api start")

	os.Remove("db.sqlite3")

	db, err := database.InitDB()
	if err != nil {
		log.Println("err open databases")
		return
	}
	defer db.Close()

	retCode := m.Run()

	os.Exit(retCode)
}

func TestApi(t *testing.T) {
	services.Init.All()

	router := gin.New()
	router.Use(gin.Logger()) // if you want to use standard logger anyway
	router.Use(mw.Sys.Recover)
	// router := gin.Default()

	//Only here success!
	router.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// router.Use(cors.Default())

	router.POST("/api/user", controllers.User.Post)
	router.GET("/api/user", controllers.User.GetAll)
	router.GET("/api/user/:id", controllers.User.Get)
	router.PUT("/api/user/:id", controllers.User.Put)
	router.DELETE("/api/user/:id", controllers.User.Delete)
	router.POST("/api/user/ufile", controllers.User.Upload)
	router.POST("/api/user/query", controllers.User.Query)
	router.GET("/api/dept", controllers.Dept.GetAll)
	router.GET("/api/dept/:id", controllers.Dept.Get)
	router.DELETE("/api/dept/:id", controllers.Dept.Delete)
	router.GET("/api/proj", controllers.Proj.GetAll)
	router.DELETE("/api/proj/:id", controllers.Proj.Delete)
	router.GET("/api/init/all", controllers.Init.All)
	router.GET("/api/init/user", controllers.Init.User)
	router.GET("/api/init/dept", controllers.Init.Dept)
	router.GET("/api/init/proj", controllers.Init.Proj)
	router.DELETE("/api/init/all", controllers.Init.Clear)
	router.DELETE("/api/init/user", controllers.Init.UserClear)
	router.DELETE("/api/init/dept", controllers.Init.DeptClear)
	router.DELETE("/api/init/proj", controllers.Init.ProjClear)
	router.GET("/api/", controllers.Home.Get)
	router.GET("/api/index", controllers.Home.Get)
	router.GET("/api/error", controllers.Home.Error)
	router.StaticFS("/static", http.Dir("static"))

	router.POST("/api/file/ufile", controllers.File.Upload)
	router.POST("/api/file/ufile2", controllers.File.Upload2)
	router.POST("/api/file/ufiles", controllers.File.Uploads)

	router.PUT("/api/auth/login", mw.Auth.LoginHandler)
	// router.NoRoute(mw.Auth.MiddlewareFunc(), mw.Sys.NoRoute)
	router.NoRoute(mw.Sys.NoRoute)

	auth := router.Group("/api/auth")

	// Refresh time can be longer than token timeout
	auth.GET("/refresh", mw.Auth.RefreshHandler)
	auth.Use(mw.Auth.MiddlewareFunc())
	{
		auth.GET("", controllers.Auth.Get)
	}

	// router.Use(cors.Default()) //error!

	router.Run(":8000")
}