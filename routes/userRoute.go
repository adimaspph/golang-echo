package routes

import (
	"golang-echo/handler"
	"golang-echo/middlewares"
)

func userRoute() {

	user := e.Group("/user")
	user.GET("", handler.GetUser)
	user.POST("", handler.AddUser)
	user.POST("/upload", handler.Upload)

	admin := e.Group("/admin", middlewares.Role("admin"))
	admin.POST("/",handler.MiddleUser)
}
