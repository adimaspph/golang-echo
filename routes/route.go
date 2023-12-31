package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func New() *echo.Echo {
	// Membuat objek echo
	e = echo.New()

	// Static
	e.Static("/static", "./template")

	// File
	e.File("/public", "./template/view/simple.html")

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	homeRoute()
	userRoute()
	cookieRoute()
	Fileserver()
	helloRoute()
	queryRoute()
	pathParameterRoute()
	headerRoute()
	downloadRoute()
	redirectRoute()
	middleRoute()
	uploadRoute()

	return e
}
