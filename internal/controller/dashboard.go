package controller

import (
	"github.com/dennis0126/network-monitor/internal/view"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DashboardController struct {
}

func NewDashboardController() DashboardController {
	return DashboardController{}
}

func (c DashboardController) RegisterRoutes(e *echo.Echo) {
	e.GET("/dashboard", c.DashboardView)
}

func (c DashboardController) DashboardView(ctx echo.Context) error {
	return render(ctx, http.StatusOK, view.Dashboard())
}
