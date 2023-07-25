package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Route
	e.GET("/hello", hello)

	// Patch Matching Order
	// Param
	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})

	// Static
	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})

	// Match any
	e.GET("/users/1/files/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/1/files/*")
	})

	// Group
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "/admin/test")
	})

	// Ruote Naming
	route := e.POST("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	route.Name = "create-user"

	// or using the inline syntax
	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}).Name = "get-user"

	// URI Building
	// Handler
	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}

	// Route
	e.GET("/users/:id", h)

	// List Router
	// Routes
	e.POST("/users", createUser)
	e.GET("/users", findUser)
	e.PUT("/users", updateUser)
	e.DELETE("/users", deleteUser)

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	os.WriteFile("./routing/routes.json", data, 0644)

	e.Logger.Fatal(e.Start("localhost:1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handlers
func createUser(c echo.Context) error {
	return nil
}

func findUser(c echo.Context) error {
	return nil
}

func updateUser(c echo.Context) error {
	return nil
}

func deleteUser(c echo.Context) error {
	return nil
}
