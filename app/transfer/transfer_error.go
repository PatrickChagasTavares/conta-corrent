package transfer

import (
	"net/http"

	"github.com/patrickchagastavares/conta-corrent/model"
)

var (
	errTransferFrom    = model.NewError(http.StatusBadRequest, "Não foi possivel encontrar sua conta", nil)
	errTransferTo      = model.NewError(http.StatusBadRequest, "Não foi possivel encontrar a conta de destino", nil)
	errTransferBalance = model.NewError(http.StatusBadRequest, "Saldo insuficiente", nil)
	errtransfer        = model.NewError(http.StatusInternalServerError, "Tivemos um problema ao transferir o valor", nil)

	errListIDNotInformed = model.NewError(http.StatusBadRequest, "Id não informado", nil)
	errListByID          = model.NewError(http.StatusInternalServerError, "Tivemos um problema ao listar as suas transferencias", nil)
)
