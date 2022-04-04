package account

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/conta-corrent/app"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
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

// list swagger document
// @Description Usado para listar todos as contas
// @Tags Account
// @Produce json
// @Success 200 {object} model.Response{Data=[]model.Account}
// @Failure 500 {object} model.Response{error=model.Error}
// @Router /v1/account [get]
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

// list swagger document
// @Description Usado para mostrar saldo da conta
// @Tags Account
// @Produce json
// @QueryParam id query int false "id da conta"
// @Success 200 {object} model.Response{Data=[]model.Account}
// @Failure 400 {object} model.Response{error=model.Error}
// @Failure 500 {object} model.Response{error=model.Error}
// @Router /v1/account/{id}/balance [get]
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

// list swagger document
// @Description Usado para mostrar saldo da conta
// @Tags Account
// @Produce json
// @Param conta body model.Account true "expected structure"
// @Success 200 {object} model.Response{Data=[]model.Account}
// @Failure 400 {object} model.Response{error=model.Error}
// @Failure 500 {object} model.Response{error=model.Error}
// @Router /v1/account [post]
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
