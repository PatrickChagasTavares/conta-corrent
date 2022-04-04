package model

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransferValidate(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr error

		InputTransfer Transfer
	}{
		"return: valid with sucess": {
			InputTransfer: Transfer{OriginID: 1, DestinationID: 2, Amount: *big.NewInt(100)},
			ExpectedErr:   nil,
		},
		"return error: id origin required": {
			InputTransfer: Transfer{DestinationID: 2, Amount: *big.NewInt(100)},
			ExpectedErr:   errTransferFromNotInput,
		},
		"return error: id destination required": {
			InputTransfer: Transfer{OriginID: 1, Amount: *big.NewInt(100)},
			ExpectedErr:   errTransferToNotInput,
		},
		"return error: value  required": {
			InputTransfer: Transfer{OriginID: 1, DestinationID: 2, Amount: *big.NewInt(0)},
			ExpectedErr:   errTransferValueNotInput,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			err := cs.InputTransfer.Validate()

			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}
