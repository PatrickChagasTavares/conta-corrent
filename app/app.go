package app

import (
	"time"

	"github.com/patrickchagastavares/conta-corrent/app/account"
	"github.com/patrickchagastavares/conta-corrent/app/login"
	"github.com/patrickchagastavares/conta-corrent/app/transfer"
	"github.com/patrickchagastavares/conta-corrent/store"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
	"github.com/patrickchagastavares/conta-corrent/utils/session"
	"github.com/patrickchagastavaresconta-corrent/utils/password"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	Account  account.App
	Login    login.App
	Transfer transfer.App
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Stores  *store.Container
	Session session.Session

	StartedAt time.Time
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	password := password.NewPassword()
	account := account.NewApp(opts.Stores, password)

	container := &Container{
		Account:  account,
		Login:    login.NewApp(opts.Stores, opts.Session, account, password),
		Transfer: transfer.NewApp(opts.Stores, account),
	}

	logger.Info("Initialized -> App")

	return container

}
