package routes

import "golang-echo/handler"

func pathParameterRoute() {
	e.GET("/get-path-param/:id", handler.PathParam)
}
