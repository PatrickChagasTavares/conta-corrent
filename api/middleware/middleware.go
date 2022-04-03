package middleware

import "github.com/patrickchagastavares/conta-corrent/utils/session"

// Options struct de opções para a criação de uma instancia dos middlewares
type Options struct {
	SessionJwt session.Session
}

// Container é um container para as implementações
type Container struct {
	Session SessionMiddleware
}

// Register middlewares for the BFF
func Register(opts Options) *Container {
	return &Container{
		Session: newSessionMiddleware(opts.SessionJwt),
	}
}
