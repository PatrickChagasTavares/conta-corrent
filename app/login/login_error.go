package login

import (
	"net/http"

	"github.com/patrickchagastavares/StoneTest/model"
)

var (
	errLoginCPFNotInput    = model.NewError(http.StatusBadRequest, "CPF não informado", nil)
	errLoginSecretNotInput = model.NewError(http.StatusBadRequest, "Secret não informado", nil)

	errLogin                = model.NewError(http.StatusBadRequest, "não foi possivel fazer seu login", nil)
	errLoginPasswordInvalid = model.NewError(http.StatusBadRequest, "senha inválida", nil)
)
