package routes

import "golang-echo/handler"

func uploadRoute() {
	upload := e.Group("/upload")
	upload.POST("/", handler.Upload)
}
