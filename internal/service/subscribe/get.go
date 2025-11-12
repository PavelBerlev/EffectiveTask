package subscribe

import (
	"EffectiveTask/internal/model"
	"context"
	"errors"
	"net/http"
)

func (s *subService) Get(ctx context.Context, id int64) (*model.SubscribeModel, int, error) {
	s.logger.Printf("Service Get: start id=%d", id)

	sub, err := s.subscribeRepo.GetSubscibeByID(ctx, id)
	if err != nil {
		s.logger.Printf("Service Get: error GetSubscibeByID id=%d err=%v", id, err)
		return nil, http.StatusInternalServerError, err
	}
	if sub == nil {
		s.logger.Printf("Service Get: subscribe not found id=%d", id)
		return nil, http.StatusNotFound, errors.New("subscribe not exist")
	}
	s.logger.Printf("Service Get: success id=%d user=%s service=%s", id, sub.UserID, sub.ServiceName)
	return sub, http.StatusOK, nil
}
