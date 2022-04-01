package health

import (
	"context"
	"time"

	"github.com/patrickchagastavares/StoneTest/utils/logger"

	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/store"
)

// App interface de health para implementação
type App interface {
	Ping(ctx context.Context) (*model.Health, error)
}

// NewApp cria uma nova instancia do serviço de health
func NewApp(stores *store.Container, startedAt time.Time) App {
	return &appImpl{
		stores:    stores,
		startedAt: startedAt,
	}
}

type appImpl struct {
	stores    *store.Container
	startedAt time.Time
}

func (s *appImpl) Ping(ctx context.Context) (health *model.Health, err error) {
	health, err = s.stores.Health.Ping(ctx)
	if err != nil {
		logger.ErrorContext(ctx, err)
		return nil, err
	}

	health.ServerStatedAt = s.startedAt.UTC().String()

	return health, nil
}
