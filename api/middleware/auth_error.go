package middleware

import (
	"net/http"

	"github.com/patrickchagastavares/conta-corrent/model"
)

var (
	errUnauthorized = model.NewError(http.StatusUnauthorized, "não autorizado", nil)
)
