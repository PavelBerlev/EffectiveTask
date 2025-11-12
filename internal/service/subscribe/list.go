package subscribe

import (
	"EffectiveTask/internal/model"
	"context"
	"net/http"
)

func (s *subService) List(ctx context.Context, limit, offset int) ([]*model.SubscribeModel, int, error) {
	s.logger.Printf("Service List: start limit=%d offset=%d", limit, offset)

	subs, err := s.subscribeRepo.ListSubscriptions(ctx, limit, offset)
	if err != nil {
		s.logger.Printf("Service List: error ListSubscriptions limit=%d offset=%d err=%v", limit, offset, err)
		return nil, http.StatusInternalServerError, err
	}
	s.logger.Printf("Service List: success limit=%d offset=%d count=%d", limit, offset, len(subs))
	return subs, http.StatusOK, nil
}
