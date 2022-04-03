package account

import (
	"net/http"

	"github.com/patrickchagastavares/StoneTest/model"
)

var (
	errAccountIDNotFound = model.NewError(http.StatusBadRequest, "Não conseguimos recuperar o id da conta informado", nil)
	errAccountCreateBind = model.NewError(http.StatusBadRequest, "Não conseguimos recuperar as informações para cria conta", nil)
)
