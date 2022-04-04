package session

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/utils/logger"
)

type Session interface {
	Generate(ctx context.Context, account *model.Account) (tokenString string, expirationTime time.Time, err error)
	LoadSession(ctx context.Context, tokenString string) (context.Context, error)
}

type sessionImpl struct {
	secret string
}

func NewSession(secret string) Session {
	session := &sessionImpl{secret}

	logger.Info("Registered -> session")

	return session
}

func (s *sessionImpl) Generate(ctx context.Context, account *model.Account) (tokenString string, expirationTime time.Time, err error) {
	expirationTime = time.Now().Add(5 * time.Minute)

	session := &SessionJWT{
		AccountOriginID: account.ID,
		Name:            account.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)

	tokenString, err = token.SignedString([]byte(s.secret))
	if err != nil {
		logger.ErrorContext(ctx, err)
		return tokenString, expirationTime, errGenerateToken
	}
	return tokenString, expirationTime, nil
}

func (s *sessionImpl) LoadSession(ctx context.Context, tokenString string) (context.Context, error) {
	token, err := jwt.ParseWithClaims(tokenString, &SessionJWT{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		logger.ErrorContext(ctx, err)
		return nil, errGetSession
	}

	if claims, ok := token.Claims.(*SessionJWT); ok && token.Valid {
		return context.WithValue(ctx, "session", *claims), nil
	}

	return nil, errTokenExpired
}

func FromContext(ctx context.Context) *SessionJWT {
	sessCtx := ctx.Value("session")
	if sessCtx == nil {
		return nil
	}
	sess := sessCtx.(SessionJWT)
	return &sess
}
