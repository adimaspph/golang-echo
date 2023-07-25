package middlewares

import (
	"golang-echo/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Role(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u := new(domain.User)
			if err := c.Bind(u); err != nil {
				return c.JSON(http.StatusBadRequest, "bad request")
			}

			allowed := false
			for _, r := range roles {
				if r == u.Role {
					allowed = true
					break
				}
			}
			if !allowed {
				return c.JSON(http.StatusUnauthorized, "You do not have permission to access this resource")
			}

			c.Set("authName", u.Name)
			return next(c)
		}
	}
}
