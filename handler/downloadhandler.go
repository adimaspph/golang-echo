package handler

import "github.com/labstack/echo/v4"

func DownloadHandler(c echo.Context) error {
	return c.Attachment("static/download.png", "static/download.png")
}
