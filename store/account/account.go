package account

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
)

// Store interface para implementação do account
type Store interface {
	List(ctx context.Context) ([]*model.Account, error)
	GetBalanceByID(ctx context.Context, id int) (*model.Account, error)
	CpfExists(ctx context.Context, cpf string) (bool, error)
	Create(ctx context.Context, account *model.Account) error
}

type storeImpl struct {
	reader *sqlx.DB
	write  *sqlx.DB
}

// NewStore cria uma nova instancia do repositorio de account
func NewStore(reader, write *sqlx.DB) Store {
	return &storeImpl{reader, write}
}

func (s *storeImpl) List(ctx context.Context) (resp []*model.Account, err error) {
	resp = make([]*model.Account, 0)

	query := `select id, name, cpf, created_at from accounts order by name`
	err = s.reader.SelectContext(ctx, &resp, query)
	if err != nil {
		logger.ErrorContext(ctx, err)
		return nil, err
	}

	return resp, nil
}

func (s *storeImpl) GetBalanceByID(ctx context.Context, id int) (resp *model.Account, err error) {
	resp = &model.Account{}

	query := `select name, cpf, balance,created_at from accounts where id = $1`
	err = s.reader.GetContext(ctx, resp, query, id)
	if err != nil {
		logger.ErrorContext(ctx, err)
		return nil, err
	}

	return resp, nil
}

func (s *storeImpl) CpfExists(ctx context.Context, cpf string) (exists bool, err error) {

	query := "SELECT EXISTS(SELECT TRUE FROM accounts WHERE cpf=$1)"
	err = s.reader.GetContext(ctx, &exists, query, cpf)
	if err != nil {
		logger.ErrorContext(ctx, err)
		return exists, err
	}

	return exists, nil
}

func (s *storeImpl) Create(ctx context.Context, account *model.Account) (err error) {
	query := `insert into accounts (name, cpf, secret_hash, secret_salt) values ($1, $2, $3, $4)`
	_, err = s.write.ExecContext(ctx, query, account.Name, account.CPF, account.SecretHash, account.SecretSalt)
	if err != nil {
		logger.ErrorContext(ctx, err)
		return err
	}

	return nil
}
