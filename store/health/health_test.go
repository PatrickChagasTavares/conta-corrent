package health

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/test"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Health

		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {
			ExpectedData: &model.Health{DatabaseStatus: "OK"},
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing()
			},
		},
		"deve retornar erro com a mensagem: ocorreu um erro": {
			ExpectedErr: model.NewError(http.StatusInternalServerError, "ocorreu um erro", nil),
			PrepareMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(errors.New("ocorreu um erro"))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := NewStore(db)
			ctx := context.Background()

			data, err := store.Ping(ctx)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}
