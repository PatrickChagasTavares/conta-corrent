package login

import (
	"context"

	"github.com/patrickchagastavares/StoneTest/app/account"
	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/store"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
	"github.com/patrickchagastavares/StoneTest/utils/password"
	"github.com/patrickchagastavares/StoneTest/utils/session"
)

type App interface {
	Login(ctx context.Context, cpf string, secret string) (*session.SessionAuth, error)
}

type appImpl struct {
	stores   *store.Container
	session  session.Session
	account  account.App
	password password.Password
}

// NewApp cria uma nova instancia do modulo login
func NewApp(stores *store.Container, session session.Session, account account.App, password password.Password) App {
	return &appImpl{
		stores:  stores,
		session: session,
		account: account,
	}
}

func (a *appImpl) Login(ctx context.Context, cpf string, secret string) (*session.SessionAuth, error) {

	if cpf == "" {
		return nil, errLoginCPFNotInput
	}

	if secret == "" {
		return nil, errLoginSecretNotInput
	}

	user := &model.Account{
		CPF:    cpf,
		Secret: secret,
	}
	if err := user.CpfIsValid(); err != nil {
		return nil, err
	}

	account, err := a.account.GetByCpf(ctx, user.CPF)
	if err != nil {
		logger.ErrorContext(ctx, err)
		return nil, errLogin
	}

	if !a.password.Verify(user.Secret, account.SecretHash, account.SecretSalt) {
		return nil, errLoginPasswordInvalid
	}

	tokenString, expirationTime, err := a.session.Generate(ctx, account)
	if err != nil {
		return nil, err
	}

	return &session.SessionAuth{
		Token:      tokenString,
		Expiration: expirationTime,
	}, nil
}
