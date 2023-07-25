package handler

import (
	"golang-echo/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var users = map[int]*domain.Users{
	1: {ID: 1, Name: "maduy", Email: "maduy@maduy.com"},
	2: {ID: 2, Name: "Erwin Smith", Email: "Erwin@smith.com"},
}

func GetUserHandlers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	user, found := users[id]
	if !found {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

// func GetUserHandlers(c echo.Context) error {
// 	name := c.Param("name")
// 	var foundUser *domain.Users
// 	for _, user := range users {
// 		if user.Name == name {
// 			foundUser = user
// 			break
// 		}
// 	}
// 	if foundUser == nil {
// 		return c.JSON(http.StatusNotFound, "User not found")
// 	}
// 	return c.JSON(http.StatusOK, foundUser)
// }
