package account

import (
	"context"

	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/store"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
	"github.com/patrickchagastavares/StoneTest/utils/password"
)

type App interface {
	List(ctx context.Context) ([]*model.Account, error)
	GetBalanceByID(ctx context.Context, id int) (*model.Account, error)
	GetByCpf(ctx context.Context, cpf string) (*model.Account, error)
	GetByID(ctx context.Context, id int) (*model.Account, error)
	Create(ctx context.Context, account *model.Account) error
	UpdateBalance(ctx context.Context, account *model.Account) error
}

type appImpl struct {
	stores *store.Container
}

// NewApp cria uma nova instancia do modulo accounts
func NewApp(stores *store.Container) App {
	return &appImpl{
		stores: stores,
	}
}

func (a *appImpl) List(ctx context.Context) ([]*model.Account, error) {
	accounts, err := a.stores.Account.List(ctx)
	if err != nil {
		return nil, errAccountList
	}

	return accounts, nil
}

func (a *appImpl) GetBalanceByID(ctx context.Context, id int) (*model.Account, error) {
	if id == 0 {
		return nil, errAccountID
	}

	account, err := a.stores.Account.GetBalanceByID(ctx, id)
	if err != nil {
		return nil, errAccountBalanceByID
	}

	return account, nil
}

func (a *appImpl) Create(ctx context.Context, account *model.Account) error {

	if err := account.Validate(); err != nil {
		return err
	}

	exists, err := a.stores.Account.CpfExists(ctx, account.CPF)
	if err != nil {
		return errAccountCreate
	}

	if exists {
		return errAccountCpfExists
	}

	account.SecretSalt = password.Salt()
	account.SecretHash = password.Encode(account.Secret, account.SecretSalt)

	if err := a.stores.Account.Create(ctx, account); err != nil {
		logger.ErrorContext(ctx, err)
		return errAccountCreate
	}

	return nil
}

func (a *appImpl) GetByCpf(ctx context.Context, cpf string) (*model.Account, error) {

	if cpf == "" {
		return nil, errAccountCpfNotInput
	}

	account, err := a.stores.Account.GetByCpf(ctx, cpf)
	if err != nil {
		return nil, errAccountGetByCpf
	}

	return account, nil
}

func (a *appImpl) GetByID(ctx context.Context, id int) (*model.Account, error) {
	if id <= 0 {
		return nil, errAccountID
	}

	account, err := a.stores.Account.GetByID(ctx, id)
	if err != nil {
		return nil, errAccountGetByCpf
	}

	return account, nil
}

func (a *appImpl) UpdateBalance(ctx context.Context, account *model.Account) error {

	if account.ID <= 0 {
		return errAccountID
	}

	if account.Balance < 0 {
		return errAccountBalance
	}

	if err := a.stores.Account.UpdateBalance(ctx, account); err != nil {
		logger.ErrorContext(ctx, err)
		return errAccountUpdateBalance
	}

	return nil
}
