package api

import (
	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/conta-corrent/api/middleware"
	v1 "github.com/patrickchagastavares/conta-corrent/api/v1"
	"github.com/patrickchagastavares/conta-corrent/app"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
)

// Options struct de opções para a criação de uma instancia das rotas
type Options struct {
	Group *echo.Group
	Apps  *app.Container

	Middleware *middleware.Container
}

// Register register routes
func Register(opts Options) {
	v1.Register(opts.Group, opts.Apps, opts.Middleware)

	logger.Info("Initialized -> Api")

}
