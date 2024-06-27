package controller

import (
	"github.com/dennis0126/network-monitor/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (c UserController) RegisterRoutes(e *echo.Echo) {
	userRouteGroup := e.Group("/users")
	userRouteGroup.GET("", c.ListUsers)
}

func (c UserController) ListUsers(ctx echo.Context) error {
	users, err := c.userService.ListUsers()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, users)
}
