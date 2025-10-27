package outbox

import (
	"context"
	"fmt"

	"github.com/darialissi/msa_big_tech/lib/postgres"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
)

func (r *Repository) SaveEvent(ctx context.Context, in *models.Event) error {
	const api = "outbox.Repository.SaveEvent"

	row, err := FromModel(in)

	if err != nil {
		return err
	}

	qb := r.sb.Insert(tableOutboxEvents).
		Columns(tableOutboxEventsColumns...).
		Values(row.Values(tableOutboxEventsColumns...)...)

	pool := r.db.GetQueryEngine(ctx)

	if _, err := pool.Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, postgres.ConvertPGError(err))
	}

	return nil
}
