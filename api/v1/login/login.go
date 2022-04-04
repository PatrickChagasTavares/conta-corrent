package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/conta-corrent/app"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
	_ "github.com/patrickchagastavares/conta-corrent/utils/session" // used swagger
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

// list swagger document
// @Description realiza login
// @Tags Login
// @Produce json
// @Param login body auth true "expected structure"
// @Success 200 {object} model.Response{Data=session.SessionAuth}
// @Failure 400 {object} model.Response{error=model.Error}
// @Failure 500 {object} model.Response{error=model.Error}
// @Router /v1/login [post]
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
