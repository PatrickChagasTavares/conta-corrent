package login

import (
	"context"

	"github.com/patrickchagastavares/conta-corrent/app/account"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
	"github.com/patrickchagastavares/conta-corrent/utils/password"
	"github.com/patrickchagastavares/conta-corrent/utils/session"
)

type App interface {
	Login(ctx context.Context, cpf string, secret string) (*session.SessionAuth, error)
}

type appImpl struct {
	session  session.Session
	account  account.App
	password password.Password
}

// NewApp cria uma nova instancia do modulo login
func NewApp(session session.Session, account account.App, password password.Password) App {
	return &appImpl{
		session:  session,
		account:  account,
		password: password,
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
