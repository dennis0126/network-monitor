package controller

import (
	"errors"
	"github.com/dennis0126/network-monitor/internal/service"
	"github.com/dennis0126/network-monitor/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{authService: authService}
}

func (c AuthController) RegisterRoutes(e *echo.Echo) {
	authRouteGroup := e.Group("/auth")
	authRouteGroup.POST("/login", c.Login)
}

type LoginDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (c AuthController) Login(ctx echo.Context) error {
	body := LoginDto{}
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	if err := ctx.Validate(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	session, err := c.authService.Login(service.LoginParam{Name: body.Name, Password: body.Password})
	if errors.Is(err, service.ErrAuthFailed) {
		return ctx.JSON(http.StatusUnauthorized, utils.NewError(err))
	}
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "SessionId"
	cookie.Value = session.ID
	cookie.Secure = true
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)
	return ctx.JSON(http.StatusOK, nil)
}
