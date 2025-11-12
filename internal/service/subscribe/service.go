package subscribe

import (
	"EffectiveTask/internal/config"
	"EffectiveTask/internal/dto"
	"EffectiveTask/internal/model"
	"EffectiveTask/internal/repository/subscribe"
	"context"
	"log"
)

type SubService interface {
	Create(ctx context.Context, req *dto.SubscribeCreateRequest) (int64, int, error)
	Update(ctx context.Context, req *dto.SubscribeUpdateRequest) (int64, int, error)
	Get(ctx context.Context, id int64) (*model.SubscribeModel, int, error)
	Delete(ctx context.Context, id int64) (int64, int, error)
	List(ctx context.Context, limit, offset int) ([]*model.SubscribeModel, int, error)
	CalculateTotalCost(ctx context.Context, req *dto.SubscribeTotalCostRequest) (int64, int, error)
}

type subService struct {
	cfg           *config.Config
	subscribeRepo subscribe.SubRepository
	logger        *log.Logger
}

func NewService(cfg *config.Config, subscribeRepo subscribe.SubRepository, logger *log.Logger) SubService {
	return &subService{
		cfg:           cfg,
		subscribeRepo: subscribeRepo,
		logger:        logger,
	}
}
