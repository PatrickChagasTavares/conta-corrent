package account

import (
	"net/http"

	"github.com/patrickchagastavares/StoneTest/model"
)

var (
	errAccountList          = model.NewError(http.StatusInternalServerError, "Tivemos um problema ao listar as contas", nil)
	errAccountBalanceByID   = model.NewError(http.StatusInternalServerError, "Tivemos um problema ao buscar o balanço da conta indicada", nil)
	errAccountID            = model.NewError(http.StatusBadRequest, "O id da conta é obrigatório", nil)
	errAccountBalance       = model.NewError(http.StatusBadRequest, "O valor do balanço não pode ser negativo", nil)
	errAccountCpfExists     = model.NewError(http.StatusBadRequest, "O cpf informado já está cadastrado", nil)
	errAccountCreate        = model.NewError(http.StatusInternalServerError, "Tivemos um problema ao criar sua conta", nil)
	errAccountGetByCpf      = model.NewError(http.StatusInternalServerError, "Tivemos um problema ao buscar sua conta", nil)
	errAccountUpdateBalance = model.NewError(http.StatusInternalServerError, "Tivemos um problema ao atualizar o saldo da conta", nil)

	errAccountCpfNotInput = model.NewError(http.StatusBadRequest, "O cpf é obrigatório", nil)
)
