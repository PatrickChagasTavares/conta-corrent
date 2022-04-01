package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/StoneTest/app"
	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
)

// Register group health check
func Register(g *echo.Group, apps *app.Container) {
	h := &handler{
		apps: apps,
	}

	g.GET("", h.ping)

	logger.Info("health Register")
}

type handler struct {
	apps *app.Container
}

func (h *handler) ping(c echo.Context) error {
	ctx := c.Request().Context()

	status, err := h.apps.Health.Ping(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: status,
	})
}
