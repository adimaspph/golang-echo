package handler

import (
	"fmt"
	"golang-echo/domain"
	"golang-echo/dto"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
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

func Upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	imgName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	path := fmt.Sprintf("./public/img/%s", imgName)
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}
