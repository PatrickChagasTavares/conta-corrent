package transfer

import (
	"context"
	"errors"
	"math/big"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/test"
	"github.com/stretchr/testify/assert"
)

func TestListByID(t *testing.T) {
	now := time.Now()
	inputId := 1

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData []*model.Transfer

		InputID     int
		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"return sucess with rows": {
			InputID: inputId,
			ExpectedData: []*model.Transfer{
				{ID: 1, OriginID: inputId, DestinationID: 2, AmountDB: "100", Amount: *big.NewInt(100), CreatedAt: now},
				{ID: 2, OriginID: inputId, DestinationID: 3, AmountDB: "2000", Amount: *big.NewInt(2000), CreatedAt: now},
			},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				rows := test.NewRows("id", "origin_id", "destination_id", "amount", "created_at").
					AddRow(1, inputId, 2, "100", now).
					AddRow(2, inputId, 3, "2000", now)

				mock.ExpectQuery("select id, origin_id, destination_id, amount, created_at  from transfers where origin_id = \\$1").
					WithArgs(inputId).
					WillReturnRows(rows)
			},
		},
		"return error": {
			InputID:     1,
			ExpectedErr: errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select id, origin_id, destination_id, amount, created_at  from transfers where origin_id = \\$1").
					WithArgs(inputId).
					WillReturnError(errors.New("an error has occurred"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := NewStore(db, db)
			ctx := context.Background()

			data, err := store.ListByID(ctx, cs.InputID)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestCreate(t *testing.T) {
	account := &model.Transfer{OriginID: 1, DestinationID: 2, Amount: *big.NewInt(100)}
	cases := map[string]struct {
		ExpectedErr error

		InputAccount *model.Transfer
		PrepareMock  func(mock sqlmock.Sqlmock)
	}{
		"return sucess with rows": {
			InputAccount: account,
			ExpectedErr:  nil,
			PrepareMock: func(mock sqlmock.Sqlmock) {

				mock.ExpectExec(
					regexp.QuoteMeta("insert into transfers (origin_id, destination_id, amount) values ($1, $2, $3)")).
					WithArgs(account.OriginID, account.DestinationID, account.Amount.String()).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"return error": {
			InputAccount: account,
			ExpectedErr:  errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(
					regexp.QuoteMeta("insert into transfers (origin_id, destination_id, amount) values ($1, $2, $3)")).
					WithArgs(account.OriginID, account.DestinationID, account.Amount.String()).
					WillReturnError(errors.New("an error has occurred"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := NewStore(db, db)
			ctx := context.Background()

			err := store.Create(ctx, cs.InputAccount)

			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}
