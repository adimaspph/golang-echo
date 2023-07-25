package routes

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func New() *echo.Echo {
	// Membuat objek echo
	e = echo.New()

	// Static
	e.Static("/static", "./template")

	// File
	e.File("/public", "./template/view/simple.html")

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

	return e
}
