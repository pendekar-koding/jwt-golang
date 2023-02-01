package main

import (
	"auth-golang/controllers"
	"auth-golang/database"
	"auth-golang/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=auth-golang port=5432 sslmode=disable"
	database.Connect(dsn)
	database.Migrate()

	router := initRouter()
	router.Run(":8000")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/auth")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
