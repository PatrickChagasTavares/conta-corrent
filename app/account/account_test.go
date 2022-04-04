package account

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/patrickchagastavares/conta-corrent/mocks"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/store"
	"github.com/patrickchagastavares/conta-corrent/test"
	"github.com/stretchr/testify/assert"
)

func TestLisByID(t *testing.T) {
	data := []*model.Account{
		{ID: 1, Name: "Account 1", Balance: *big.NewInt(100), CreatedAt: time.Now()},
		{ID: 2, Name: "Account 2", Balance: *big.NewInt(100), CreatedAt: time.Now()},
	}

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData []*model.Account

		PrepareMock func(mock *mocks.MockAccountStore)
	}{
		"return sucess": {
			ExpectedData: data,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().List(gomock.Any()).Return(data, nil)
			},
		},
		"return erro: store": {
			ExpectedErr: errAccountList,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().List(gomock.Any()).Return(nil, errors.New("error"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mock := mocks.NewMockAccountStore(ctrl)

			cs.PrepareMock(mock)

			app := NewApp(&store.Container{Account: mock}, nil)

			data, err := app.List(ctx)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestGetBalanceByID(t *testing.T) {
	data := &model.Account{ID: 1, Name: "Account 1", Balance: *big.NewInt(100), CreatedAt: time.Now()}
	id := 1

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Account

		InputID     int
		PrepareMock func(mock *mocks.MockAccountStore)
	}{
		"return sucess": {
			InputID:      id,
			ExpectedData: data,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetBalanceByID(gomock.Any(), id).Return(data, nil)
			},
		},
		"return erro: id required": {
			InputID:     -1,
			ExpectedErr: errAccountID,
			PrepareMock: func(mock *mocks.MockAccountStore) {},
		},
		"return erro: store": {
			InputID:     id,
			ExpectedErr: errAccountBalanceByID,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetBalanceByID(gomock.Any(), id).Return(nil, errors.New("error"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mock := mocks.NewMockAccountStore(ctrl)

			cs.PrepareMock(mock)

			app := NewApp(&store.Container{Account: mock}, nil)

			data, err := app.GetBalanceByID(ctx, cs.InputID)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestGetByCpf(t *testing.T) {
	cpf := "12345678901"
	data := &model.Account{ID: 1, Name: "Account 1", CPF: cpf, SecretHash: "secret_hash", SecretSalt: "secret_salt", Balance: *big.NewInt(100)}

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Account

		InputCPF    string
		PrepareMock func(mock *mocks.MockAccountStore)
	}{
		"return sucess": {
			InputCPF:     cpf,
			ExpectedData: data,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByCpf(gomock.Any(), cpf).Return(data, nil)
			},
		},
		"return erro: id required": {
			InputCPF:    "",
			ExpectedErr: errAccountCpfNotInput,
			PrepareMock: func(mock *mocks.MockAccountStore) {},
		},
		"return erro: store": {
			InputCPF:    cpf,
			ExpectedErr: errAccountGet,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByCpf(gomock.Any(), cpf).Return(nil, errors.New("error"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mock := mocks.NewMockAccountStore(ctrl)

			cs.PrepareMock(mock)

			app := NewApp(&store.Container{Account: mock}, nil)

			data, err := app.GetByCpf(ctx, cs.InputCPF)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestGetByID(t *testing.T) {
	id := 1
	data := &model.Account{ID: id, Balance: *big.NewInt(100)}

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Account

		InputID     int
		PrepareMock func(mock *mocks.MockAccountStore)
	}{
		"return sucess": {
			InputID:      id,
			ExpectedData: data,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByID(gomock.Any(), id).Return(data, nil)
			},
		},
		"return erro: id required": {
			InputID:     -1,
			ExpectedErr: errAccountID,
			PrepareMock: func(mock *mocks.MockAccountStore) {},
		},
		"return erro: store": {
			InputID:     id,
			ExpectedErr: errAccountGet,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByID(gomock.Any(), id).Return(nil, errors.New("error"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mock := mocks.NewMockAccountStore(ctrl)

			cs.PrepareMock(mock)

			app := NewApp(&store.Container{Account: mock}, nil)

			data, err := app.GetByID(ctx, cs.InputID)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestUpdateBalance(t *testing.T) {
	value := big.NewInt(1000)
	fmt.Println(value.Sub(value, big.NewInt(1100)))
	cases := map[string]struct {
		ExpectedErr error

		Inputaccount *model.Account
		PrepareMock  func(mock *mocks.MockAccountStore)
	}{
		"return sucess": {
			Inputaccount: &model.Account{ID: 1, Balance: *big.NewInt(100)},
			ExpectedErr:  nil,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().UpdateBalance(gomock.Any(), &model.Account{ID: 1, Balance: *big.NewInt(100)}).Return(nil)
			},
		},
		"return erro: id required": {
			Inputaccount: &model.Account{ID: -1},
			ExpectedErr:  errAccountID,
			PrepareMock:  func(mock *mocks.MockAccountStore) {},
		},
		"return erro: store": {
			Inputaccount: &model.Account{ID: 1, Balance: *big.NewInt(100)},
			ExpectedErr:  errAccountUpdateBalance,
			PrepareMock: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().UpdateBalance(gomock.Any(), &model.Account{ID: 1, Balance: *big.NewInt(100)}).Return(errors.New("error"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mock := mocks.NewMockAccountStore(ctrl)

			cs.PrepareMock(mock)

			app := NewApp(&store.Container{Account: mock}, nil)

			err := app.UpdateBalance(ctx, cs.Inputaccount)

			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestCreate(t *testing.T) {
	cpfSucess := "74596687021"
	inputCorrect := &model.Account{Name: "teste", CPF: cpfSucess, Secret: "secret"}

	cases := map[string]struct {
		ExpectedErr error

		Inputaccount        *model.Account
		PrepareMockAccount  func(mock *mocks.MockAccountStore)
		PrepareMockPassword func(mock *mocks.MockPassword)
	}{
		"return sucess": {
			Inputaccount: inputCorrect,
			ExpectedErr:  nil,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().CpfExists(gomock.Any(), cpfSucess).Return(false, nil)
				mock.EXPECT().Create(gomock.Any(), inputCorrect).Return(nil)
			},
			PrepareMockPassword: func(mock *mocks.MockPassword) {
				mock.EXPECT().Salt().Return("salt")
				mock.EXPECT().Encode("secret", "salt").Return("hash")
			},
		},
		"return erro: name required": {
			Inputaccount:        &model.Account{CPF: cpfSucess, Secret: "secret"},
			ExpectedErr:         model.NewError(http.StatusBadRequest, "O nome é obrigatório", nil),
			PrepareMockAccount:  func(mock *mocks.MockAccountStore) {},
			PrepareMockPassword: func(mock *mocks.MockPassword) {},
		},
		"return erro: store cpf exists": {
			Inputaccount: inputCorrect,
			ExpectedErr:  errAccountCreate,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().CpfExists(gomock.Any(), cpfSucess).Return(false, errors.New("error"))
			},
			PrepareMockPassword: func(mock *mocks.MockPassword) {},
		},
		"return erro: exists cpf": {
			Inputaccount: inputCorrect,
			ExpectedErr:  errAccountCpfExists,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().CpfExists(gomock.Any(), cpfSucess).Return(true, nil)
			},
			PrepareMockPassword: func(mock *mocks.MockPassword) {},
		},
		"return erro: store create": {
			Inputaccount: inputCorrect,
			ExpectedErr:  errAccountCreate,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().CpfExists(gomock.Any(), cpfSucess).Return(false, nil)
				mock.EXPECT().Create(gomock.Any(), inputCorrect).Return(errors.New("error"))
			},
			PrepareMockPassword: func(mock *mocks.MockPassword) {
				mock.EXPECT().Salt().Return("salt")
				mock.EXPECT().Encode("secret", "salt").Return("hash")
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockAccount := mocks.NewMockAccountStore(ctrl)
			mockPassword := mocks.NewMockPassword(ctrl)

			cs.PrepareMockAccount(mockAccount)
			cs.PrepareMockPassword(mockPassword)

			app := NewApp(&store.Container{Account: mockAccount}, mockPassword)

			err := app.Create(ctx, cs.Inputaccount)

			assert.Equal(t, err, cs.ExpectedErr)
		})
	}
}
