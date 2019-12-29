package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Router is manage URL endpoint.
func Router(e *echo.Echo) {
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Word!")
	})
}