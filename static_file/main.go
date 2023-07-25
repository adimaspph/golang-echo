package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// base_url/static/test.html
	e.Static("/static", "./static_file/template")

	// base_url/P1011894.jpg
	e.Static("/", "./upload/img")

	e.Logger.Fatal(e.Start("localhost:1323"))
}
