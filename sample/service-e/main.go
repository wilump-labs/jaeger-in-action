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
	e.GET("/enter", enter)
	e.GET("/exit", exit)
	e.GET("/monitor", monitor)

	// Start server
	e.Logger.Fatal(e.Start(":1500"))
}

func enter(c echo.Context) error {
	return c.String(http.StatusOK, "enter")
}

func exit(c echo.Context) error {
	return c.String(http.StatusOK, "exit")
}

func monitor(c echo.Context) error {
	return c.String(http.StatusOK, "monitor")
}
