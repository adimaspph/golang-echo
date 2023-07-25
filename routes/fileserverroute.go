package routes

import "github.com/labstack/echo/v4"

func Fileserver() {
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return c.File("static/index.html")
	})

}
