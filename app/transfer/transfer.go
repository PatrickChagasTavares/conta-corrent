package transfer

import (
	"context"

	"github.com/patrickchagastavares/conta-corrent/app/account"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/store"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
	"golang.org/x/sync/errgroup"
)

type App interface {
	Create(ctx context.Context, transfer *model.Transfer) error
	ListByID(ctx context.Context, accountID int) ([]*model.Transfer, error)
}

type appImpl struct {
	stores  *store.Container
	account account.App
}

// NewApp cria uma nova instancia do modulo login
func NewApp(stores *store.Container, account account.App) App {
	return &appImpl{
		stores:  stores,
		account: account,
	}
}

func (a *appImpl) Create(ctx context.Context, transfer *model.Transfer) error {
	var (
		accountFrom *model.Account
		accountTo   *model.Account
		errs        errgroup.Group
	)

	if err := transfer.Validate(); err != nil {
		return err
	}

	errs.Go(func() (err error) {
		accountFrom, err = a.account.GetByID(ctx, transfer.OriginID)
		if err != nil {
			logger.ErrorContext(ctx, err)
			return errTransferFrom
		}
		return nil
	})

	errs.Go(func() (err error) {
		accountTo, err = a.account.GetByID(ctx, transfer.DestinationID)
		if err != nil {
			logger.ErrorContext(ctx, err)
			return errTransferTo
		}

		return nil
	})

	if err := errs.Wait(); err != nil {
		return err
	}

	if accountFrom.Balance.CmpAbs(&transfer.Amount) < 0 {
		return errTransferBalance
	}

	accountFrom.Balance.Sub(&accountFrom.Balance, &transfer.Amount)
	accountTo.Balance.Add(&accountTo.Balance, &transfer.Amount)

	if err := a.executeTransfer(ctx, &errs, accountFrom, accountTo); err != nil {
		accountFrom.Balance.Add(&accountFrom.Balance, &transfer.Amount)
		accountTo.Balance.Sub(&accountTo.Balance, &transfer.Amount)

		go a.executeTransfer(context.Background(), &errs, accountFrom, accountTo)
		logger.ErrorContext(ctx, err)
		return err
	}

	if err := a.stores.Transfer.Create(ctx, transfer); err != nil {
		return errtransfer
	}

	return nil
}

func (a *appImpl) ListByID(ctx context.Context, accountID int) ([]*model.Transfer, error) {

	if accountID <= 0 {
		return nil, errListIDNotInformed
	}

	transfers, err := a.stores.Transfer.ListByID(ctx, accountID)
	if err != nil {
		logger.ErrorContext(ctx, err)
		return nil, errListByID
	}

	return transfers, nil
}

func (a *appImpl) executeTransfer(ctx context.Context, errs *errgroup.Group, accountFrom, accountTo *model.Account) error {
	errs.Go(func() (err error) {
		if err := a.account.UpdateBalance(context.Background(), accountFrom); err != nil {
			logger.ErrorContext(ctx, err)
			return errtransfer
		}

		return nil
	})

	errs.Go(func() (err error) {
		if err := a.account.UpdateBalance(ctx, accountTo); err != nil {
			logger.ErrorContext(ctx, err)
			return errtransfer
		}

		return nil
	})

	if err := errs.Wait(); err != nil {
		return err
	}

	return nil
}
