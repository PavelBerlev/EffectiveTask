package subscribe

import (
	"EffectiveTask/internal/dto"
	"EffectiveTask/internal/model"
	"context"
	"errors"
	"net/http"
)

func (s *subService) Update(ctx context.Context, req *dto.SubscribeUpdateRequest) (int64, int, error) {
	s.logger.Printf("Service Update: start id=%d user=%s service=%s", req.ID, req.UserID, req.ServiceName)

	//check if exist
	subExist, err := s.subscribeRepo.GetSubscibeByID(ctx, req.ID)

	if err != nil {
		s.logger.Printf("Service Update: error GetSubscibeByID id=%d err=%v", req.ID, err)
		return 0, http.StatusInternalServerError, err
	}

	if subExist == nil {
		s.logger.Printf("Service Update: subscribe not found id=%d", req.ID)
		return 0, http.StatusBadRequest, errors.New("subscribe not exist")
	}

	//validate date
	start, end, err := dto.ValidateDateRange(req.StartDate, req.EndDate)
	if err != nil {
		s.logger.Printf("Service Update: date validation error: %v", err)
		return 0, http.StatusBadRequest, err
	}

	//Create subscribe record
	subModel := &model.SubscribeModel{
		ID:          req.ID,
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   start,
		EndDate:     end,
	}

	id, err := s.subscribeRepo.UpdateSubscribe(ctx, subModel)
	if err != nil {
		s.logger.Printf("Service Update: error UpdateSubscribe id=%d err=%v", req.ID, err)
		return 0, http.StatusInternalServerError, err
	}

	s.logger.Printf("Service Update: success id=%d user=%s service=%s price=%d", id, req.UserID, req.ServiceName, req.Price)
	return id, http.StatusCreated, nil
}
