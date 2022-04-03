package login

import (
	"net/http"

	"github.com/patrickchagastavares/conta-corrent/model"
)

var (
	errLoginBind = model.NewError(http.StatusBadRequest, "corpo da requisão é invalido", nil)
)
