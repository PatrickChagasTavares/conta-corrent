package transfer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/conta-corrent/api/middleware"
	"github.com/patrickchagastavares/conta-corrent/app"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
	"github.com/patrickchagastavares/conta-corrent/utils/session"
)

type handler struct {
	apps *app.Container
}

func Register(g *echo.Group, apps *app.Container, middleware *middleware.Container) {
	h := &handler{
		apps: apps,
	}

	g.GET("", h.list, middleware.Session.Private)
	g.POST("", h.create, middleware.Session.Private)

	logger.Info("transfer Register")
}

// list swagger document
// @Description Usado para listar todos as transferências
// @Tags Transfer
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response{Data=[]model.Transfer}
// @Failure 500 {object} model.Response{error=model.Error}
// @Router /v1/transfer [get]
func (h *handler) list(c echo.Context) error {
	ctx := c.Request().Context()

	sess := session.FromContext(ctx)

	transfers, err := h.apps.Transfer.ListByID(ctx, sess.AccountOriginID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: transfers,
	})
}

// list swagger document
// @Description Usado para realizar/criar uma transferencia
// @Tags Transfer
// @Produce json
// @Security ApiKeyAuth
// @Param transferencia body model.Transfer true "expected structure"
// @Success 204
// @Failure 400 {object} model.Response{error=model.Error}
// @Failure 500 {object} model.Response{error=model.Error}
// @Router /v1/transfer [post]
func (h *handler) create(c echo.Context) error {
	ctx := c.Request().Context()

	sess := session.FromContext(ctx)

	transfer := new(model.Transfer)
	if err := c.Bind(transfer); err != nil {
		return errTransferBind
	}

	transfer.OriginID = sess.AccountOriginID

	if err := h.apps.Transfer.Create(ctx, transfer); err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}
