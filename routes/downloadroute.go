package routes

import (
	"golang-echo/handler"
)

func downloadRoute() {
	e.GET("/download", handler.DownloadHandler)
}
