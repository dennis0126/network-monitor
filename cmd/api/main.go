package main

import (
	"context"
	"github.com/dennis0126/network-monitor/internal/config"
	"github.com/dennis0126/network-monitor/internal/controller"
	"github.com/dennis0126/network-monitor/internal/db"
	"github.com/dennis0126/network-monitor/internal/repository"
	"github.com/dennis0126/network-monitor/internal/service"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
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

	cfg := config.InitConfig()

	// database connection
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.DbString)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer conn.Close(ctx)
	queries := db.New(conn)

	// repository
	userRepository := repository.NewUserRepository(ctx, queries)

	// service
	userService := service.NewUserService(userRepository)

	// controller
	userController := controller.NewUserController(userService)

	userController.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(cfg.Port)))
}
