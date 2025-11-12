package subscribe

import (
	"EffectiveTask/internal/model"
	"context"
	"database/sql"
)

func (r *subRepository) GetSubscibeByID(ctx context.Context, id int64) (*model.SubscribeModel, error) {
	query := `
        SELECT id, user_id, service_name, price, start_date, end_date, created_at, updated_at
        FROM subscriptions
        WHERE id = $1 AND deleted_at IS NULL
    `

	r.logger.Printf("Repo GetSubscibeByID: start id=%d", id)

	row := r.db.QueryRowContext(ctx, query, id)

	sub := &model.SubscribeModel{}
	err := row.Scan(
		&sub.ID,
		&sub.UserID,
		&sub.ServiceName,
		&sub.Price,
		&sub.StartDate,
		&sub.EndDate,
		&sub.CreatedAt,
		&sub.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Printf("Repo GetSubscibeByID: not found id=%d", id)
			return nil, nil
		}
		r.logger.Printf("Repo GetSubscibeByID: error id=%d err=%v", id, err)
		return nil, err
	}

	r.logger.Printf("Repo GetSubscibeByID: success id=%d user=%s service=%s", id, sub.UserID, sub.ServiceName)
	return sub, nil
}
