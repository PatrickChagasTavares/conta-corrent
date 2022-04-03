package v1

import (
	"github.com/patrickchagastavares/conta-corrent/api/middleware"
	"github.com/patrickchagastavares/conta-corrent/api/v1/account"
	"github.com/patrickchagastavares/conta-corrent/api/v1/login"
	"github.com/patrickchagastavares/conta-corrent/api/v1/transfer"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/conta-corrent/app"
)

// Register regristra as rotas v1
func Register(g *echo.Group, apps *app.Container, middleware *middleware.Container) {
	v1 := g.Group("/v1")

	account.Register(v1.Group("/account"), apps)
	login.Register(v1.Group("/login"), apps)
	transfer.Register(v1.Group("/transfer"), apps, middleware)

}
