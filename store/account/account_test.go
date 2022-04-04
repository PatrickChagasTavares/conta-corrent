package account

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

func TestList(t *testing.T) {
	now := time.Now()
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData []*model.Account

		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"return sucess": {
			ExpectedData: []*model.Account{
				{ID: 1, Name: "Teste", CPF: "12345678901", CreatedAt: now},
				{ID: 2, Name: "Teste2", CPF: "12345678902", CreatedAt: now},
			},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				rows := test.NewRows("id", "name", "cpf", "created_at").
					AddRow(1, "Teste", "12345678901", now).
					AddRow(2, "Teste2", "12345678902", now)

				mock.ExpectQuery("select id, name, cpf, created_at from accounts order by name").
					WillReturnRows(rows)
			},
		},
		"return error": {
			ExpectedErr: errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select id, name, cpf, created_at from accounts order by name").
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

			data, err := store.List(ctx)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestGetBalanceByID(t *testing.T) {
	now := time.Now()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Account

		InputID     int
		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"return sucess with rows": {
			InputID:      1,
			ExpectedData: &model.Account{Name: "Teste", CPF: "12345678901", BalanceDB: "10000", Balance: *big.NewInt(10000), CreatedAt: now},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				rows := test.NewRows("name", "cpf", "balance", "created_at").
					AddRow("Teste", "12345678901", "10000", now)

				mock.ExpectQuery("select name, cpf, balance, created_at from accounts where id = \\$1").
					WithArgs(1).
					WillReturnRows(rows)
			},
		},
		"return error": {
			InputID:     1,
			ExpectedErr: errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select name, cpf, balance, created_at from accounts where id = \\$1").
					WithArgs(1).
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

			data, err := store.GetBalanceByID(ctx, cs.InputID)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestGetByCpf(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Account

		InputCPF    string
		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"return sucess with rows": {
			InputCPF:     "12345678901",
			ExpectedData: &model.Account{ID: 1, Name: "Teste", CPF: "12345678901", SecretHash: "teste_hash", SecretSalt: "teste_salt", BalanceDB: "10000", Balance: *big.NewInt(10000)},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				rows := test.NewRows("id", "name", "cpf", "secret_hash", "secret_salt", "balance").
					AddRow(1, "Teste", "12345678901", "teste_hash", "teste_salt", "10000")

				mock.ExpectQuery("select id, name, cpf, secret_hash, secret_salt, balance from accounts where cpf = \\$1").
					WithArgs("12345678901").
					WillReturnRows(rows)
			},
		},
		"return error": {
			InputCPF:    "12345678901",
			ExpectedErr: errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select id, name, cpf, secret_hash, secret_salt, balance from accounts where cpf = \\$1").
					WithArgs("12345678901").
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

			data, err := store.GetByCpf(ctx, cs.InputCPF)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestGetByID(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Account

		InputID     int
		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"return sucess with rows": {
			InputID:      1,
			ExpectedData: &model.Account{ID: 1, BalanceDB: "10000", Balance: *big.NewInt(10000)},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				rows := test.NewRows("id", "balance").
					AddRow(1, "10000")

				mock.ExpectQuery("select id, balance from accounts where id = \\$1").
					WithArgs(1).
					WillReturnRows(rows)
			},
		},
		"return error": {
			InputID:     1,
			ExpectedErr: errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select id, balance from accounts where id = \\$1").
					WithArgs(1).
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

			data, err := store.GetByID(ctx, cs.InputID)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestCpfExists(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData bool

		InputCPF    string
		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"return sucess with rows": {
			InputCPF:     "12345678901",
			ExpectedData: true,
			PrepareMock: func(mock sqlmock.Sqlmock) {
				rows := test.NewRows("exists").
					AddRow(true)

				mock.ExpectQuery(
					regexp.QuoteMeta("select exists(SELECT TRUE FROM accounts WHERE cpf= $1)"),
				).
					WithArgs("12345678901").
					WillReturnRows(rows)
			},
		},
		"return error": {
			InputCPF:    "12345678901",
			ExpectedErr: errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta("select exists(SELECT TRUE FROM accounts WHERE cpf= $1)"),
				).
					WithArgs("12345678901").
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

			data, err := store.CpfExists(ctx, cs.InputCPF)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestCreate(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr error

		InputAccount *model.Account
		PrepareMock  func(mock sqlmock.Sqlmock)
	}{
		"return sucess with rows": {
			InputAccount: &model.Account{Name: "teste", CPF: "12345678901", SecretHash: "teste_hash", SecretSalt: "teste_salt"},
			ExpectedErr:  nil,
			PrepareMock: func(mock sqlmock.Sqlmock) {

				mock.ExpectExec(
					regexp.QuoteMeta("insert into accounts (name, cpf, secret_hash, secret_salt) values ($1, $2, $3, $4)")).
					WithArgs("teste", "12345678901", "teste_hash", "teste_salt").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"return error": {
			InputAccount: &model.Account{Name: "teste", CPF: "12345678901", SecretHash: "teste_hash", SecretSalt: "teste_salt"},
			ExpectedErr:  errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(
					regexp.QuoteMeta("insert into accounts (name, cpf, secret_hash, secret_salt) values ($1, $2, $3, $4)")).
					WithArgs("teste", "12345678901", "teste_hash", "teste_salt").
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

func TestUpdateBalance(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr error

		InputAccount *model.Account
		PrepareMock  func(mock sqlmock.Sqlmock)
	}{
		"return sucess with rows": {
			InputAccount: &model.Account{ID: 1, Balance: *big.NewInt(10000)},
			ExpectedErr:  nil,
			PrepareMock: func(mock sqlmock.Sqlmock) {

				mock.ExpectExec(
					regexp.QuoteMeta("update accounts set balance=$1, updated_at=CURRENT_TIMESTAMP where id=$3")).
					WithArgs("10000", 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		"return error": {
			InputAccount: &model.Account{ID: 1, Balance: *big.NewInt(10000)},
			ExpectedErr:  errors.New("an error has occurred"),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(
					regexp.QuoteMeta("update accounts set balance=$1, updated_at=CURRENT_TIMESTAMP where id=$3")).
					WithArgs("10000", 1).
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

			err := store.UpdateBalance(ctx, cs.InputAccount)

			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}
