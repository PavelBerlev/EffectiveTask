package subscribe

import (
	"EffectiveTask/internal/dto"
	"EffectiveTask/internal/model"
	"context"
	"errors"
	"net/http"
)

func (s *subService) Create(ctx context.Context, req *dto.SubscribeCreateRequest) (int64, int, error) {
	s.logger.Printf("Service Create: start user=%s service=%s start_date=%s", req.UserID, req.ServiceName, req.StartDate)

	//check if subscribe record already exists
	subExist, err := s.subscribeRepo.GetSubscribeByUserIDAndServiceName(ctx, req.UserID, req.ServiceName)

	if err != nil {
		s.logger.Printf("Service Create: error GetSubscribeByUserIDAndServiceName: %v", err)
		return 0, http.StatusInternalServerError, err
	}

	if subExist != nil {
		s.logger.Printf("Service Create: already exists user=%s service=%s id=%d", req.UserID, req.ServiceName, subExist.ID)
		return 0, http.StatusBadRequest, errors.New("subscribe already exist")
	}

	//validate date must be like 02-2006, end_date > start_date
	start, end, err := dto.ValidateDateRange(req.StartDate, req.EndDate)
	if err != nil {
		s.logger.Printf("Service Create: date validation error: %v", err)
		return 0, http.StatusBadRequest, err
	}

	//Create subscribe record
	subModel := &model.SubscribeModel{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   start,
		EndDate:     end,
	}

	id, err := s.subscribeRepo.CreateSubscribe(ctx, subModel)
	if err != nil {
		s.logger.Printf("Service Create: error CreateSubscribe: %v", err)
		return 0, http.StatusInternalServerError, err
	}

	s.logger.Printf("Service Create: success id=%d user=%s service=%s price=%d", id, req.UserID, req.ServiceName, req.Price)
	return id, http.StatusCreated, nil
}
