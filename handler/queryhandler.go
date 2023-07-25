package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")

	message := "Data pengguna dengan nama: " + name + " dan umur: " + age

	return c.String(http.StatusOK, message)
}
