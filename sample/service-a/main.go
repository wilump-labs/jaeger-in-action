package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/call-a", callA)
	e.GET("/call-b", callB)
	e.GET("/call-c", callC)
	e.GET("/call-f", callF)

	// Start server
	e.Logger.Fatal(e.Start(":1500"))
}

func callA(c echo.Context) error {
	return c.String(http.StatusOK, "call A")
}

func callB(c echo.Context) error {
	// HTTP call

	return c.String(http.StatusOK, "call A")
}

func callC(c echo.Context) error {
	// HTTP call

	return c.String(http.StatusOK, "call A")
}

func callF(c echo.Context) error {
	// gRPC call

	// TODO: gRPC call

	return c.String(http.StatusOK, "call A")
}
