package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HeaderHandler(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		return c.String(http.StatusBadRequest, "tidak ada Authorization header")
	}
	return c.String(http.StatusOK, ("Authhoir param = " + header))
}
