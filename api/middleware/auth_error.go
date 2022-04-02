package middleware

import (
	"net/http"

	"github.com/patrickchagastavares/StoneTest/model"
)

var (
	errUnauthorized = model.NewError(http.StatusUnauthorized, "não autorizado", nil)
)
