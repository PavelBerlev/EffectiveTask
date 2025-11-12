package subscribe

import (
	"EffectiveTask/internal/model"
	"context"
)

func (r *subRepository) CreateSubscribe(ctx context.Context, model *model.SubscribeModel) (int64, error) {
	query := `INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at`

	r.logger.Printf("Repo CreateSubscribe: start service=%s price=%d user=%s", model.ServiceName, model.Price, model.UserID)

	err := r.db.QueryRowContext(ctx, query,
		model.ServiceName,
		model.Price,
		model.UserID,
		model.StartDate,
		model.EndDate,
	).Scan(&model.ID, &model.CreatedAt, &model.UpdatedAt)

	if err != nil {
		r.logger.Printf("Repo CreateSubscribe: error: %v", err)
		return 0, err
	}

	r.logger.Printf("Repo CreateSubscribe: success id=%d", model.ID)
	return model.ID, nil
}
