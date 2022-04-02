package api

import (
	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/StoneTest/api/middleware"
	v1 "github.com/patrickchagastavares/StoneTest/api/v1"
	"github.com/patrickchagastavares/StoneTest/app"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
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
