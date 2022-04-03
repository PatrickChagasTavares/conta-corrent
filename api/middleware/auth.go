package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/conta-corrent/utils/session"
)

// SessionMiddleware it's an interface to validate some user roles
type SessionMiddleware interface {
	Public(next echo.HandlerFunc) echo.HandlerFunc
	Private(next echo.HandlerFunc) echo.HandlerFunc
}

type middlewareAuthImpl struct {
	session session.Session
}

func newSessionMiddleware(session session.Session) SessionMiddleware {
	return &middlewareAuthImpl{
		session: session,
	}
}

func (m *middlewareAuthImpl) Public(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func (m *middlewareAuthImpl) Private(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return errUnauthorized
		}

		splitted := strings.Split(authorization, " ")
		if len(splitted) != 2 {
			return errUnauthorized
		}

		ctx := c.Request().Context()

		ctx, err := m.session.LoadSession(ctx, splitted[1])
		if err != nil {
			return err
		}

		c.SetRequest(
			c.Request().WithContext(ctx),
		)

		return next(c)
	}
}
