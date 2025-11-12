package subscribe

import (
	"EffectiveTask/internal/model"
	"context"
)

func (r *subRepository) ListSubscriptions(ctx context.Context, limit, offset int) ([]*model.SubscribeModel, error) {
	query := `
        SELECT id, user_id, service_name, price, start_date, end_date, created_at, updated_at
        FROM subscriptions
        WHERE deleted_at IS NULL
        ORDER BY id DESC
        LIMIT $1 OFFSET $2
    `

	r.logger.Printf("Repo ListSubscriptions: start limit=%d offset=%d", limit, offset)

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		r.logger.Printf("Repo ListSubscriptions: error limit=%d offset=%d err=%v", limit, offset, err)
		return nil, err
	}
	defer rows.Close()

	var subs []*model.SubscribeModel
	for rows.Next() {
		sub := &model.SubscribeModel{}
		if err := rows.Scan(
			&sub.ID,
			&sub.UserID,
			&sub.ServiceName,
			&sub.Price,
			&sub.StartDate,
			&sub.EndDate,
			&sub.CreatedAt,
			&sub.UpdatedAt,
		); err != nil {
			r.logger.Printf("Repo ListSubscriptions: scan error: %v", err)
			return nil, err
		}
		subs = append(subs, sub)
	}

	if err := rows.Err(); err != nil {
		r.logger.Printf("Repo ListSubscriptions: rows error: %v", err)
		return nil, err
	}

	r.logger.Printf("Repo ListSubscriptions: success limit=%d offset=%d count=%d", limit, offset, len(subs))
	return subs, nil
}
