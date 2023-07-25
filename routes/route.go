package routes

import (
	"github.com/labstack/echo/v4"
	"golang-echo/handler"
	"net/http"
)

func New() *echo.Echo {
	// Membuat objek echo
	e := echo.New()

	// Routing sederhana
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halo, Selamat datang di Echo!")
	})

	user := e.Group("/user")
	user.GET("", handler.GetUser)

	return e
}
