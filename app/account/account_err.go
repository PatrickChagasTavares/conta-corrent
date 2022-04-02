package account

import (
	"net/http"

	"github.com/patrickchagastavares/StoneTest/model"
)

var (
	errAccountList          = model.NewError(http.StatusInternalServerError, "tivemos um problema ao listar as contas.", nil)
	errAccountBalanceByID   = model.NewError(http.StatusInternalServerError, "tivemos um problema ao buscar o balanço da conta indicada.", nil)
	errAccountID            = model.NewError(http.StatusBadRequest, "o id da conta é obrigatório.", nil)
	errAccountCpfExists     = model.NewError(http.StatusBadRequest, "o cpf informado já está cadastrado.", nil)
	errAccountCreate        = model.NewError(http.StatusInternalServerError, "tivemos um problema ao criar sua conta.", nil)
	errAccountGetByCpf      = model.NewError(http.StatusInternalServerError, "tivemos um problema ao buscar sua conta.", nil)
)
