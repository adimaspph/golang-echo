package routes

import "golang-echo/handler"

func userRoute() {

	user := e.Group("/user")
	user.GET("", handler.GetUser)
	user.POST("", handler.AddUser)
	user.POST("/upload", handler.Upload)
}
