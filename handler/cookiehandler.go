package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetCookieHandler(c echo.Context) error {
	cookie := &http.Cookie{
		Name:  "username",
		Value: "john_doe",
		Path:  "/",
	}
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "Cookie set successfully")
}

func GetCookieHandler(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return c.String(http.StatusNotFound, "Cookie not found")
	}

	return c.String(http.StatusOK, "Username: "+cookie.Value)
}
