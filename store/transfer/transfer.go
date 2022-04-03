package transfer

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
)

// Store interface para implementação do transfer
type Store interface {
	Create(ctx context.Context, transfer *model.Transfer) error
	ListByID(ctx context.Context, id int) ([]*model.Transfer, error)
}

type storeImpl struct {
	reader *sqlx.DB
	write  *sqlx.DB
}

// NewStore cria uma nova instancia do repositorio de transfer
func NewStore(reader, write *sqlx.DB) Store {
	return &storeImpl{reader, write}
}

func (s *storeImpl) Create(ctx context.Context, transfer *model.Transfer) error {
	query := `insert into transfers (origin_id, destination_id, amount) values ($1, $2, $3)`
	_, err := s.write.ExecContext(ctx, query, transfer.OriginID, transfer.DestinationID, transfer.Value.String())
	if err != nil {
		logger.ErrorContext(ctx, err)
		return err
	}

	return nil
}

func (s *storeImpl) ListByID(ctx context.Context, id int) (transfers []*model.Transfer, err error) {
	transfers = make([]*model.Transfer, 0)

	query := `select id, origin_id, destination_id, amount, created_at  from transfers where origin_id = $1`
	err = s.reader.SelectContext(ctx, &transfers, query, id)
	if err != nil {
		logger.ErrorContext(ctx, err)
		return nil, err
	}

	return transfers, nil
}
