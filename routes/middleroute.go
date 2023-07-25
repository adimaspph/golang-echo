package routes

import (
	"golang-echo/handler"
	"golang-echo/middlewares"
)

func middleRoute() {
	admin := e.Group("/admin", middlewares.Role("admin"))
	admin.POST("/", handler.MiddleUser)
}
