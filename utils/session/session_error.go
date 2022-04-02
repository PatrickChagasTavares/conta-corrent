package session

import (
	"net/http"

	"github.com/patrickchagastavares/StoneTest/model"
)

var (
	errGenerateToken = model.NewError(http.StatusInternalServerError, "Tivemos um problema para gerar o seu token", nil)
	errGetSession    = model.NewError(http.StatusInternalServerError, "Tivemos um problema para recuperar sua sessão", nil)
	errTokenExpired  = model.NewError(http.StatusBadRequest, "Seu token expirou", nil)
)
