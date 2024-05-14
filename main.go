package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
