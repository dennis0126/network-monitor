package controller

import (
	"errors"
	"github.com/dennis0126/network-monitor/internal/service"
	"github.com/dennis0126/network-monitor/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type AuthController struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthController(authService service.AuthService, userService service.UserService) AuthController {
	return AuthController{authService: authService, userService: userService}
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

	cookie := http.Cookie{
		Name:     "SessionId",
		Value:    session.ID,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	ctx.SetCookie(&cookie)
	return ctx.JSON(http.StatusOK, nil)
}

type SessionAuthConfig struct {
	Skipper middleware.Skipper
}

var DefaultSessionAuthConfig = SessionAuthConfig{
	Skipper: middleware.DefaultSkipper,
}

func (c AuthController) SessionAuth(config SessionAuthConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultSessionAuthConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if config.Skipper(ctx) {
				return next(ctx)
			}

			sessionId, err := ctx.Cookie("SessionId")
			if err != nil {
				return ctx.Redirect(http.StatusTemporaryRedirect, "/")
			}
			session, err := c.authService.GetSessionById(sessionId.Value)
			if err != nil {
				return err
			}
			if session == nil {
				return ctx.Redirect(http.StatusTemporaryRedirect, "/")
			}

			user, err := c.userService.GetUserById(session.UserID)
			ctx.Set("user", user)

			return next(ctx)
		}
	}
}
