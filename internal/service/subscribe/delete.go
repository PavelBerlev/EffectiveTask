package subscribe

import (
	"context"
	"errors"
	"net/http"
)

func (s *subService) Delete(ctx context.Context, id int64) (int64, int, error) {
	s.logger.Printf("Service Delete: start id=%d", id)

	sub, err := s.subscribeRepo.GetSubscibeByID(ctx, id)
	if err != nil {
		s.logger.Printf("Service Delete: error GetSubscibeByID id=%d err=%v", id, err)
		return 0, http.StatusInternalServerError, err
	}
	if sub == nil {
		s.logger.Printf("Service Delete: subscribe not found id=%d", id)
		return 0, http.StatusNotFound, errors.New("subscribe not exist")
	}

	deletedID, err := s.subscribeRepo.DeleteSubscribe(ctx, id)
	if err != nil {
		s.logger.Printf("Service Delete: error DeleteSubscribe id=%d err=%v", id, err)
		return 0, http.StatusInternalServerError, err
	}

	s.logger.Printf("Service Delete: success id=%d", deletedID)
	return deletedID, http.StatusOK, nil
}
