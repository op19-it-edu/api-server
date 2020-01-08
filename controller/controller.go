package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Hello, Worldを返す関数
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
