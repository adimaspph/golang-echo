package routes

import "golang-echo/handler"

func queryRoute() {
	e.GET("/api/users", handler.GetUsersHandler)
}
