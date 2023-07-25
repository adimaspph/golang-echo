package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func homeRoute() {
	// Routing sederhana
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Halo, Selamat datang di Echo!")
	})
}
