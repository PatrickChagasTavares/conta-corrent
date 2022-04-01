package app

import (
	"time"

	"github.com/patrickchagastavares/StoneTest/app/account"
	"github.com/patrickchagastavares/StoneTest/app/health"
	"github.com/patrickchagastavares/StoneTest/store"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	Health  health.App
	Account account.App
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Stores *store.Container

	StartedAt time.Time
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	container := &Container{
		Health:  health.NewApp(opts.Stores, opts.StartedAt),
		Account: account.NewApp(opts.Stores),
	}

	logger.Info("Initialized -> App")

	return container

}
