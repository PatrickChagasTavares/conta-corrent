package account

import (
	"net/http"
	"strconv"

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

	g.GET("", h.list)
	g.POST("", h.create)
	g.GET("/:id/balance", h.balance)

	logger.Info("accounts Register")
}

func (h *handler) list(c echo.Context) error {
	ctx := c.Request().Context()

	accounts, err := h.apps.Account.List(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: accounts,
	})
}

func (h *handler) balance(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.ErrorContext(ctx, err)
		return errAccountIDNotFound
	}

	account, err := h.apps.Account.GetBalanceByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: account,
	})
}

func (h *handler) create(c echo.Context) error {
	ctx := c.Request().Context()

	account := new(model.Account)
	if err := c.Bind(account); err != nil {
		logger.ErrorContext(ctx, err)
		return errAccountCreateBind
	}

	if err := h.apps.Account.Create(ctx, account); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{
		Data: account,
	})
}
