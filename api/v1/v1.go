package v1

import (
	"github.com/patrickchagastavares/StoneTest/api/middleware"
	"github.com/patrickchagastavares/StoneTest/api/v1/account"
	"github.com/patrickchagastavares/StoneTest/api/v1/health"
	"github.com/patrickchagastavares/StoneTest/api/v1/login"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/StoneTest/app"
)

// Register regristra as rotas v1
func Register(g *echo.Group, apps *app.Container, middleware *middleware.Container) {
	v1 := g.Group("/v1")

	health.Register(v1.Group("/health"), apps)
	account.Register(v1.Group("/account"), apps)
	login.Register(v1.Group("/login"), apps)

}
