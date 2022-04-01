package v1

import (
	"github.com/patrickchagastavares/StoneTest/api/v1/health"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/StoneTest/app"
)

// Register regristra as rotas v1
func Register(g *echo.Group, apps *app.Container) {
	v1 := g.Group("/v1")

	health.Register(v1.Group("/health"), apps)

}
