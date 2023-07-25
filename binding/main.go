package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type UserDTO struct {
	Name    string
	Email   string
	IsAdmin bool
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		// Load into separate struct for security
		user := UserDTO{
			Name:    u.Name,
			Email:   u.Email,
			IsAdmin: false, // avoids exposing field that should not be bound
		}

		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start("localhost:1323"))
}
