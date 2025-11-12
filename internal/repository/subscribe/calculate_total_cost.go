package subscribe

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func (r *subRepository) CalculateTotalCost(ctx context.Context, startDate, endDate time.Time, userID *uuid.UUID, serviceName *string) (int64, error) {
	query := `
        SELECT COALESCE(SUM(price), 0)
        FROM subscriptions
        WHERE deleted_at IS NULL
        AND start_date <= $1
        AND (end_date IS NULL OR end_date >= $2)
    `

	args := []interface{}{endDate, startDate}

	if userID != nil {
		query += ` AND user_id = $3`
		args = append(args, *userID)
		r.logger.Printf("Repo CalculateTotalCost: start start_date=%s end_date=%s user_id=%s", startDate, endDate, userID.String())
	} else if serviceName != nil {
		query += ` AND service_name = $3`
		args = append(args, *serviceName)
		r.logger.Printf("Repo CalculateTotalCost: start start_date=%s end_date=%s service_name=%s", startDate, endDate, *serviceName)
	}

	var totalCost int64
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&totalCost)
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Printf("Repo CalculateTotalCost: no rows found")
			return 0, nil
		}
		r.logger.Printf("Repo CalculateTotalCost: error: %v", err)
		return 0, err
	}

	r.logger.Printf("Repo CalculateTotalCost: success total_cost=%d", totalCost)
	return totalCost, nil
}
