package model

import (
	"math/big"
	"net/http"
	"time"
)

var (
	errTransferFromNotInput  = NewError(http.StatusBadRequest, "O id da sua conta é obrigatório", nil)
	errTransferToNotInput    = NewError(http.StatusBadRequest, "O id do destinatário é obrigatório", nil)
	errTransferValueNotInput = NewError(http.StatusBadRequest, "O valor informado deve ser maior que zero", nil)
)

type Transfer struct {
	ID            int       `json:"id" db:"id"`
	OriginID      int       `json:"account_origin_id" db:"origin_id"`
	DestinationID int       `json:"account_destination_id" db:"destination_id"`
	AmountDB      string    `json:"-" db:"amount"`
	Amount        big.Int   `json:"amount" db:"-"`
	CreateAt      time.Time `json:"created_at" db:"created_at"`
}

func (t *Transfer) Validate() error {
	if t.OriginID <= 0 {
		return errTransferFromNotInput
	}

	if t.DestinationID <= 0 {
		return errTransferToNotInput
	}

	if t.Amount.Cmp(big.NewInt(0)) <= 0 {
		return errTransferValueNotInput
	}

	return nil
}

func (t *Transfer) ConvertBigInt() {
	t.Amount.SetString(t.AmountDB, 10)
}
