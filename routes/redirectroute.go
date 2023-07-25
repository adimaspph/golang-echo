package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func redirectRoute() {
	e.GET("/tes", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://example.com")
	})
}

// e.GET("/", func(c echo.Context) error {
// 		redirectURL := c.QueryParam("redirect_ur")
// 		if redirectURL != "" {
// 			return c.Redirect(http.StatusFound, redirectURL)
// 		}
// 		return c.String(http.StatusOK, "Ini adalah halaman sekarang!")
// 	})

// 	e.GET("/otherpage", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Ini adalah halaman lain!")
// 	})
