package transfer

import (
	"net/http"

	"github.com/patrickchagastavares/conta-corrent/model"
)

var (
	errTransferBind = model.NewError(http.StatusBadRequest, "corpo da requisão é invalido", nil)
)
