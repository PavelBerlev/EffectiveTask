package subscribe

import (
	"context"
	"database/sql"
)

func (r *subRepository) DeleteSubscribe(ctx context.Context, id int64) (int64, error) {
	query := `
        UPDATE subscriptions
        SET deleted_at = NOW(), updated_at = NOW()
        WHERE id = $1 AND deleted_at IS NULL
        RETURNING id
    `

	r.logger.Printf("Repo DeleteSubscribe: start id=%d", id)

	var ID int64
	err := r.db.QueryRowContext(ctx, query, id).Scan(&ID)
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Printf("Repo DeleteSubscribe: not found id=%d", id)
			return 0, nil
		}
		r.logger.Printf("Repo DeleteSubscribe: error id=%d err=%v", id, err)
		return 0, err
	}

	r.logger.Printf("Repo DeleteSubscribe: success id=%d", ID)
	return ID, nil
}
