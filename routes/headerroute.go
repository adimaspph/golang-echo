package routes

import "golang-echo/handler"

func headerRoute() {
	e.GET("/get-header", handler.HeaderHandler)
}
