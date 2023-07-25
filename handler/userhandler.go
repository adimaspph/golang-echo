package handler

import (
	"golang-echo/domain"
	"golang-echo/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UpdateUser struct{}
)

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "get all user")
}

func AddUser(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Load into separate struct for security
	user := dto.UserDTO{
		Name:    u.Name,
		Email:   u.Email,
		IsAdmin: false, // avoids exposing field that should not be bound
	}

	return c.JSON(http.StatusOK, user)
}
