package routes

import "github.com/labstack/echo/v4"

func Fileserver() {
	e.Static("/static", "static")
 maduy
	e.GET("/file", func(c echo.Context) error {

	e.GET("/file-server", func(c echo.Context) error {
main
		return c.File("static/index.html")
	})

}
