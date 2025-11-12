package subscribe

import (
	"EffectiveTask/internal/model"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func (r *subRepository) GetSubscribeByUserIDAndServiceName(ctx context.Context, userId uuid.UUID, serviceName string) (*model.SubscribeModel, error) {
	query := `SELECT id, service_name, price, user_id, start_date, end_date, created_at, updated_at
	FROM subscriptions
	WHERE user_id = $1 AND service_name = $2 AND deleted_at IS NULL
	LIMIT 1`

	r.logger.Printf("Repo GetSubscribeByUserIDAndServiceName: start user=%s service=%s", userId, serviceName)

	row := r.db.QueryRowContext(ctx, query, userId, serviceName)

	var result model.SubscribeModel
	err := row.Scan(&result.ID,
		&result.ServiceName,
		&result.Price,
		&result.UserID,
		&result.StartDate,
		&result.EndDate,
		&result.CreatedAt,
		&result.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Printf("Repo GetSubscribeByUserIDAndServiceName: not found user=%s service=%s", userId, serviceName)
			return nil, nil
		}
		r.logger.Printf("Repo GetSubscribeByUserIDAndServiceName: error: %v", err)
		return nil, err
	}

	r.logger.Printf("Repo GetSubscribeByUserIDAndServiceName: success id=%d user=%s service=%s", result.ID, userId, serviceName)
	return &result, nil
}
