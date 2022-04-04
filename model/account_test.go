package model

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountValidate(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr error

		InputAccount Account
	}{
		"return: valid with sucess": {
			InputAccount: Account{Name: "test", Secret: "1234567890", CPF: "23208950004", Balance: *big.NewInt(100)},
			ExpectedErr:  nil,
		},
		"return error: name required": {
			InputAccount: Account{Secret: "1234567890", CPF: "23208950004", Balance: *big.NewInt(100)},
			ExpectedErr:  errNameRequired,
		},
		"return error: secret required": {
			InputAccount: Account{Name: "test", CPF: "23208950004", Balance: *big.NewInt(100)},
			ExpectedErr:  errSecretRequired,
		},
		"return error: cpf required": {
			InputAccount: Account{Name: "test", Secret: "1234567890", Balance: *big.NewInt(100)},
			ExpectedErr:  errCPFRequired,
		},
		"return error: cpf size invalid": {
			InputAccount: Account{Name: "test", Secret: "1234567890", CPF: "232089500041", Balance: *big.NewInt(100)},
			ExpectedErr:  errCPFSizeInvalid,
		},
		"return error: cpf invalid number unic": {
			InputAccount: Account{Name: "test", Secret: "1234567890", CPF: "00000000000", Balance: *big.NewInt(100)},
			ExpectedErr:  errCPFInvalid,
		},
		"return error: cpf invalid": {
			InputAccount: Account{Name: "test", Secret: "1234567890", CPF: "23208950002", Balance: *big.NewInt(100)},
			ExpectedErr:  errCPFInvalid,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			err := cs.InputAccount.Validate()

			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}
