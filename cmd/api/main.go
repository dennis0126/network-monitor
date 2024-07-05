package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/dennis0126/network-monitor/internal/config"
	"github.com/dennis0126/network-monitor/internal/controller"
	"github.com/dennis0126/network-monitor/internal/db"
	"github.com/dennis0126/network-monitor/internal/repository"
	"github.com/dennis0126/network-monitor/internal/service"
	"github.com/dennis0126/network-monitor/internal/view"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	cfg := config.InitConfig()

	e.Validator = controller.NewValidator()

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
	sessionRepository := repository.NewSessionRepository(ctx, queries)

	// service
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(sessionRepository, userService)
	// controller
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)

	userController.RegisterRoutes(e)
	authController.RegisterRoutes(e)

	e.GET("/", echo.WrapHandler(templ.Handler(view.Index())))

	fmt.Println(getRoutesAsString(e))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(cfg.Port)))
}

func getRoutesAsString(e *echo.Echo) string {
	var result string
	for _, route := range e.Routes() {
		result += fmt.Sprintf("%s %s -> %s\n", route.Method, route.Path, route.Name)
	}
	return result
}
