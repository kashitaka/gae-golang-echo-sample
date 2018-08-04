package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func init() {

	g := e.Group("/sample")

	g.GET("", helloMsg)
}

func helloMsg(c echo.Context) error {
	return c.JSON(http.StatusOK, "It works")
}
