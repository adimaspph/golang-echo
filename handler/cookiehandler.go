package handler

import (
	"encoding/json"
	"golang-echo/domain"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"
)

func SetCookieHandler(c echo.Context) error {
	// Data user dalam bentuk struct
	user := domain.User{
		Name:  "asep",
		Email: "asep@asep.com",
		Role:  "admin",
	}

	cookieValue, err := json.Marshal(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error marshalling user data")
	}
	//  nilai cookie agar aman untuk disimpan dalam cookie
	cookieValues := url.QueryEscape(string(cookieValue))
	cookie := &http.Cookie{
		Name:     "tezzz",
		Value:    cookieValues,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "Cookie set successfully")
}

func GetCookieHandler(c echo.Context) error {
	cookie, err := c.Cookie("tezzz")
	if err != nil {
		return c.String(http.StatusNotFound, "Cookie not found")
	}
	cookieValue, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error unescaping cookie value")
	}

	var user domain.User
	err = json.Unmarshal([]byte(cookieValue), &user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error unmarshalling user data")
	}

	return c.JSON(http.StatusOK, user)
}
