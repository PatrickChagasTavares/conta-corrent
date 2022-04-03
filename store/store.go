package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/patrickchagastavares/conta-corrent/store/account"
	"github.com/patrickchagastavares/conta-corrent/store/transfer"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
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
		Account:  account.NewStore(opts.Reader, opts.Writer),
		Transfer: transfer.NewStore(opts.Reader, opts.Writer),
	}

	logger.Info("Registered -> Store")

	return container
}
