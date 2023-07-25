package routes

import "golang-echo/handler"

func queryRoute() {
	// e.GET("/users", handler.GetUsersHandler)
	e.GET("/users/:id", handler.GetUserHandlers)
}
