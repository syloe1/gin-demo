package router

import (
	"gin-demo/handler"
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFile("/", "./static/index.html")
	api := r.Group("/api") // 路由分组🔥
	{
		//public
		api.POST("/login", handler.Login)
		api.GET("/user/:id", handler.GetUser)
		api.POST("/user", handler.CreateUser)
		api.GET("/users", handler.GetUserList)
		//private
		auth := api.Group("/")
		auth.Use(middleware.JWTAuth())
		auth.PUT("/user/:id", handler.UpdateUser)
		auth.DELETE("/user/:id", handler.DeleteUser)
	}

	return r
}