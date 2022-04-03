package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/patrickchagastavares/StoneTest/store/account"
	"github.com/patrickchagastavares/StoneTest/store/health"
	"github.com/patrickchagastavares/StoneTest/store/transfer"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	Health   health.Store
	Account  account.Store
	Transfer transfer.Store
}

// Options struct de opções para a criação de uma instancia dos repositórios
type Options struct {
	Writer *sqlx.DB
	Reader *sqlx.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	container := &Container{
		Health:   health.NewStore(opts.Reader),
		Account:  account.NewStore(opts.Reader, opts.Writer),
		Transfer: transfer.NewStore(opts.Reader, opts.Writer),
	}

	logger.Info("Registered -> Store")

	return container
}
