package routes

import (
	"golang-echo/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// Membuat objek echo
	e := echo.New()

	// Routing sederhana
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halo, Selamat datang di Echo!")
	})

	user := e.Group("/user")
	user.GET("/", handler.GetUser)
	user.POST("/", handler.AddUser)

	return e
}
