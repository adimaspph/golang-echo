package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MiddleUser(c echo.Context) error {
	// Get Context
	Name := c.Get("authName").(string)
	result := fmt.Sprintf("Hello %s", Name)

	return c.JSON(http.StatusOK, result)
}
