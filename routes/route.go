package routes

import (
	"golang-echo/handler"
	"net/http"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func New() *echo.Echo {
	// Membuat objek echo
	e = echo.New()

	homeRoute()
	userRoute()

	return e
}
