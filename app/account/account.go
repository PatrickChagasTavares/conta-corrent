package account

import (
	"context"
	"math/big"

	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/store/account"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
	"github.com/patrickchagastavares/conta-corrent/utils/password"
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
	store    account.Store
	password password.Password
}

// NewApp cria uma nova instancia do modulo accounts
func NewApp(store account.Store, password password.Password) App {
	return &appImpl{
		store:    store,
		password: password,
	}
}

func (a *appImpl) List(ctx context.Context) ([]*model.Account, error) {
	accounts, err := a.store.List(ctx)
	if err != nil {
		return nil, errAccountList
	}

	return accounts, nil
}

func (a *appImpl) GetBalanceByID(ctx context.Context, id int) (*model.Account, error) {
	if id <= 0 {
		return nil, errAccountID
	}

	account, err := a.store.GetBalanceByID(ctx, id)
	if err != nil {
		return nil, errAccountBalanceByID
	}

	return account, nil
}

func (a *appImpl) Create(ctx context.Context, account *model.Account) error {

	if err := account.Validate(); err != nil {
		return err
	}

	exists, err := a.store.CpfExists(ctx, account.CPF)
	if err != nil {
		return errAccountCreate
	}

	if exists {
		return errAccountCpfExists
	}

	account.SecretSalt = a.password.Salt()
	account.SecretHash = a.password.Encode(account.Secret, account.SecretSalt)

	if err := a.store.Create(ctx, account); err != nil {
		logger.ErrorContext(ctx, err)
		return errAccountCreate
	}

	return nil
}

func (a *appImpl) GetByCpf(ctx context.Context, cpf string) (*model.Account, error) {

	if cpf == "" {
		return nil, errAccountCpfNotInput
	}

	account, err := a.store.GetByCpf(ctx, cpf)
	if err != nil {
		return nil, errAccountGet
	}

	return account, nil
}

func (a *appImpl) GetByID(ctx context.Context, id int) (*model.Account, error) {
	if id <= 0 {
		return nil, errAccountID
	}

	account, err := a.store.GetByID(ctx, id)
	if err != nil {
		return nil, errAccountGet
	}

	return account, nil
}

func (a *appImpl) UpdateBalance(ctx context.Context, account *model.Account) error {

	if account.ID <= 0 {
		return errAccountID
	}

	if account.Balance.CmpAbs(big.NewInt(0)) < 0 {
		return errAccountBalance
	}

	if err := a.store.UpdateBalance(ctx, account); err != nil {
		logger.ErrorContext(ctx, err)
		return errAccountUpdateBalance
	}

	return nil
}
