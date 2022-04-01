package health

import (
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/patrickchagastavares/StoneTest/mocks"
	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/store"
	"github.com/patrickchagastavares/StoneTest/test"
)

func TestPing(t *testing.T) {
	startedAt := time.Now()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Health

		InputDatetime time.Time

		PrepareMock func(mock *mocks.MockHealthStore)
	}{
		"deve retornar sucesso": {
			InputDatetime: startedAt,
			ExpectedData:  &model.Health{DatabaseStatus: "OK", ServerStatedAt: startedAt.UTC().String()},
			PrepareMock: func(mock *mocks.MockHealthStore) {
				mock.EXPECT().Ping(gomock.Any()).Times(1).
					Return(&model.Health{DatabaseStatus: "OK"}, nil)
			},
		},
		"deve retornar erro com a mensagem: ocorreu um erro": {
			InputDatetime: startedAt,
			ExpectedErr:   model.NewError(http.StatusInternalServerError, "ocorreu um erro", nil),
			PrepareMock: func(mock *mocks.MockHealthStore) {
				mock.EXPECT().Ping(gomock.Any()).Times(1).
					Return(nil, model.NewError(http.StatusInternalServerError, "ocorreu um erro", nil))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)
			mock := mocks.NewMockHealthStore(ctrl)

			cs.PrepareMock(mock)

			app := NewApp(&store.Container{Health: mock}, cs.InputDatetime)

			data, err := app.Ping(ctx)

			assert.Equal(t, data, cs.ExpectedData)
			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}
