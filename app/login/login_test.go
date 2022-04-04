package login

import (
	"errors"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/patrickchagastavares/conta-corrent/mocks"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/test"
	"github.com/patrickchagastavares/conta-corrent/utils/session"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	cpfSucess := "74596687021"
	secretSucess := "secret"
	time := time.Now().Add(5 * time.Minute)

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *session.SessionAuth

		InputCPF    string
		InputSecret string

		PrepareMockAccount  func(mock *mocks.MockAccountStore)
		PrepareMockSession  func(mock *mocks.MockSession)
		PrepareMockPassword func(mock *mocks.MockPassword)
	}{
		"return sucess": {
			ExpectedData: &session.SessionAuth{
				Token:      "token",
				Expiration: time,
			},
			InputCPF:    cpfSucess,
			InputSecret: secretSucess,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByCpf(gomock.Any(), cpfSucess).Return(&model.Account{
					ID: 1, Name: "test", CPF: cpfSucess, SecretHash: "secret_hash", SecretSalt: "secret_salt", Balance: *big.NewInt(100),
				}, nil)
			},
			PrepareMockPassword: func(mock *mocks.MockPassword) {
				mock.EXPECT().Verify(secretSucess, "secret_hash", "secret_salt").Return(true)
			},
			PrepareMockSession: func(mock *mocks.MockSession) {
				mock.EXPECT().Generate(gomock.Any(), &model.Account{
					ID: 1, Name: "test", CPF: cpfSucess, SecretHash: "secret_hash", SecretSalt: "secret_salt", Balance: *big.NewInt(100),
				}).Return("token", time, nil)
			},
		},
		"return error: CPF required": {
			ExpectedErr:         errLoginCPFNotInput,
			InputSecret:         secretSucess,
			PrepareMockAccount:  func(mock *mocks.MockAccountStore) {},
			PrepareMockPassword: func(mock *mocks.MockPassword) {},
			PrepareMockSession:  func(mock *mocks.MockSession) {},
		},
		"return error: secret required": {
			ExpectedErr:         errLoginSecretNotInput,
			InputCPF:            cpfSucess,
			PrepareMockAccount:  func(mock *mocks.MockAccountStore) {},
			PrepareMockPassword: func(mock *mocks.MockPassword) {},
			PrepareMockSession:  func(mock *mocks.MockSession) {},
		},
		"return error: cpf invalid": {
			ExpectedErr:         model.NewError(http.StatusBadRequest, "O cpf é inválido", nil),
			InputCPF:            "11111111111",
			InputSecret:         secretSucess,
			PrepareMockAccount:  func(mock *mocks.MockAccountStore) {},
			PrepareMockPassword: func(mock *mocks.MockPassword) {},
			PrepareMockSession:  func(mock *mocks.MockSession) {},
		},
		"return error: store cpf not found": {
			ExpectedErr: errLogin,
			InputCPF:    cpfSucess,
			InputSecret: secretSucess,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByCpf(gomock.Any(), cpfSucess).Return(nil, errors.New("cpf not found"))
			},
			PrepareMockPassword: func(mock *mocks.MockPassword) {},
			PrepareMockSession:  func(mock *mocks.MockSession) {},
		},
		"return error: valid password": {
			ExpectedErr: errLoginPasswordInvalid,
			InputCPF:    cpfSucess,
			InputSecret: secretSucess,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByCpf(gomock.Any(), cpfSucess).Return(&model.Account{
					ID: 1, Name: "test", CPF: cpfSucess, SecretHash: "secret_hash", SecretSalt: "secret_salt", Balance: *big.NewInt(100),
				}, nil)
			},
			PrepareMockPassword: func(mock *mocks.MockPassword) {
				mock.EXPECT().Verify(secretSucess, "secret_hash", "secret_salt").Return(false)
			},
			PrepareMockSession: func(mock *mocks.MockSession) {},
		},
		"return error: generate token": {
			ExpectedErr: errors.New("error generate token"),
			InputCPF:    cpfSucess,
			InputSecret: secretSucess,
			PrepareMockAccount: func(mock *mocks.MockAccountStore) {
				mock.EXPECT().GetByCpf(gomock.Any(), cpfSucess).Return(&model.Account{
					ID: 1, Name: "test", CPF: cpfSucess, SecretHash: "secret_hash", SecretSalt: "secret_salt", Balance: *big.NewInt(100),
				}, nil)
			},
			PrepareMockPassword: func(mock *mocks.MockPassword) {
				mock.EXPECT().Verify(secretSucess, "secret_hash", "secret_salt").Return(true)
			},
			PrepareMockSession: func(mock *mocks.MockSession) {
				mock.EXPECT().Generate(gomock.Any(), &model.Account{
					ID: 1, Name: "test", CPF: cpfSucess, SecretHash: "secret_hash", SecretSalt: "secret_salt", Balance: *big.NewInt(100),
				}).Return("", time, errors.New("error generate token"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mockAccount := mocks.NewMockAccountStore(ctrl)
			mockSession := mocks.NewMockSession(ctrl)
			mockPassword := mocks.NewMockPassword(ctrl)

			cs.PrepareMockAccount(mockAccount)
			cs.PrepareMockSession(mockSession)
			cs.PrepareMockPassword(mockPassword)

			app := NewApp(mockSession, mockAccount, mockPassword)

			data, err := app.Login(ctx, cs.InputCPF, cs.InputSecret)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}
