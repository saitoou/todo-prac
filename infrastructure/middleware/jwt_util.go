package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func ExtractEmailFromToken(c echo.Context) (string, error) {
	email, ok := c.Get("user").(string)
	if !ok || email == "" {
		return "", errors.New("failed to extract token from context")
	}

	return email, nil
}
