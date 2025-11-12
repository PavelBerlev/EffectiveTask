package subscribe

import (
	"EffectiveTask/internal/model"
	"context"
	"database/sql"
)

func (r *subRepository) UpdateSubscribe(ctx context.Context, sub *model.SubscribeModel) (int64, error) {
	query := `
        UPDATE subscriptions
        SET service_name = $1, price = $2, start_date = $3, end_date = $4, updated_at = NOW()
        WHERE id = $5
        RETURNING id
    `

	r.logger.Printf("Repo UpdateSubscribe: start id=%d service=%s price=%d", sub.ID, sub.ServiceName, sub.Price)

	var id int64
	err := r.db.QueryRowContext(ctx, query,
		sub.ServiceName,
		sub.Price,
		sub.StartDate,
		sub.EndDate,
		sub.ID,
	).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Printf("Repo UpdateSubscribe: not found id=%d", sub.ID)
			return 0, nil
		}
		r.logger.Printf("Repo UpdateSubscribe: error id=%d err=%v", sub.ID, err)
		return 0, err
	}

	r.logger.Printf("Repo UpdateSubscribe: success id=%d", id)
	return id, nil
}
