package app

import (
	"time"

	"github.com/patrickchagastavares/StoneTest/app/account"
	"github.com/patrickchagastavares/StoneTest/app/health"
	"github.com/patrickchagastavares/StoneTest/app/login"
	"github.com/patrickchagastavares/StoneTest/store"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
	"github.com/patrickchagastavares/StoneTest/utils/session"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	Health  health.App
	Account account.App
	Login   login.App
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Stores  *store.Container
	Session session.Session

	StartedAt time.Time
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	account := account.NewApp(opts.Stores)

	container := &Container{
		Health:  health.NewApp(opts.Stores, opts.StartedAt),
		Account: account,
		Login:   login.NewApp(opts.Stores, opts.Session, account),
	}

	logger.Info("Initialized -> App")

	return container

}
