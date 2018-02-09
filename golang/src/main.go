package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// переопределяем контекст для данных
type CustomContext struct {
	echo.Context
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Helllo, World!\n")
	})

	e.GET("/metrics", promhttp.Handler())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
