package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func PathParam(c echo.Context) error {
	// User ID from path `/get-path-param/:param`
	param := c.Param("id")
	if param == "" {
		return c.String(http.StatusBadRequest, "tidak ada param")
	}
	return c.String(http.StatusOK, ("path param = " + param))
}
