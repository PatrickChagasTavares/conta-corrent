package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/StoneTest/app"
	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
)

type handler struct {
	apps *app.Container
}

// Register group account
func Register(g *echo.Group, apps *app.Container) {
	h := &handler{
		apps: apps,
	}

	g.POST("", h.login)

	logger.Info("login Register")
}

func (h *handler) login(c echo.Context) error {
	ctx := c.Request().Context()

	auth := new(auth)
	if err := c.Bind(auth); err != nil {
		logger.ErrorContext(ctx, err)
		return errLoginBind
	}

	login, err := h.apps.Login.Login(ctx, auth.CPF, auth.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: login,
	})
}
