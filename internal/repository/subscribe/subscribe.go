package subscribe

import (
	"EffectiveTask/internal/model"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type SubRepository interface {
	GetSubscribeByUserIDAndServiceName(ctx context.Context, userId uuid.UUID, serviceName string) (*model.SubscribeModel, error)
	GetSubscibeByID(ctx context.Context, id int64) (*model.SubscribeModel, error)
	CreateSubscribe(ctx context.Context, model *model.SubscribeModel) (int64, error)
	UpdateSubscribe(ctx context.Context, model *model.SubscribeModel) (int64, error)
	DeleteSubscribe(ctx context.Context, id int64) (int64, error)
	ListSubscriptions(ctx context.Context, limit, offset int) ([]*model.SubscribeModel, error)
	CalculateTotalCost(ctx context.Context, startDate, endDate time.Time, userID *uuid.UUID, serviceName *string) (int64, error) // <- добавлено
}

type subRepository struct {
	db     *sql.DB
	logger *log.Logger
}

func NewRepository(db *sql.DB, logger *log.Logger) SubRepository {
	return &subRepository{
		db:     db,
		logger: logger,
	}
}
