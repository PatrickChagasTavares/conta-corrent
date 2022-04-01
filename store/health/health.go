package health

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/patrickchagastavares/StoneTest/model"
)

// Store interface para implementação do health
type Store interface {
	Ping(ctx context.Context) (health *model.Health, err error)
}

// NewStore cria uma nova instancia do repositorio de health
func NewStore(reader *sqlx.DB) Store {
	return &storeImpl{reader}
}

type storeImpl struct {
	reader *sqlx.DB
}

// Ping checa se o banco está online
func (r *storeImpl) Ping(ctx context.Context) (health *model.Health, err error) {
	err = r.reader.PingContext(ctx)
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, err.Error(), nil)
	}
	return &model.Health{DatabaseStatus: "OK"}, nil
}
