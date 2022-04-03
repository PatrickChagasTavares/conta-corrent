package session

import (
	"net/http"

	"github.com/patrickchagastavares/conta-corrent/model"
)

var (
	errGenerateToken = model.NewError(http.StatusInternalServerError, "Tivemos um problema para gerar o seu token", nil)
	errGetSession    = model.NewError(http.StatusUnauthorized, "Não foi possivel recuperar a sessão informada", nil)
	errTokenExpired  = model.NewError(http.StatusUnauthorized, "Seu token expirou", nil)
)
