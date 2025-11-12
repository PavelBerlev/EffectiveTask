package subscribe

import (
	"EffectiveTask/internal/dto"
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func (s *subService) CalculateTotalCost(ctx context.Context, req *dto.SubscribeTotalCostRequest) (int64, int, error) {
	s.logger.Printf("Service CalculateTotalCost: start start_date=%s end_date=%s user_id=%v service_name=%v", req.StartDate, req.EndDate, req.UserID, req.ServiceName)

	// date check
	start, end, err := dto.ValidateDateRange(req.StartDate, &req.EndDate)
	if err != nil {
		s.logger.Printf("Service CalculateTotalCost: date validation error: %v", err)
		return 0, http.StatusBadRequest, err
	}

	// must be userId or ServiceName
	if req.UserID == nil && req.ServiceName == nil {
		s.logger.Printf("Service CalculateTotalCost: neither user_id nor service_name provided")
		return 0, http.StatusBadRequest, errors.New("either user_id or service_name must be provided")
	}

	var userID *uuid.UUID
	if req.UserID != nil {
		id, err := uuid.Parse(*req.UserID)
		if err != nil {
			s.logger.Printf("Service CalculateTotalCost: invalid user_id format: %v", err)
			return 0, http.StatusBadRequest, errors.New("invalid user_id format")
		}
		userID = &id
	}

	totalCost, err := s.subscribeRepo.CalculateTotalCost(ctx, start, *end, userID, req.ServiceName)
	if err != nil {
		s.logger.Printf("Service CalculateTotalCost: error CalculateTotalCost err=%v", err)
		return 0, http.StatusInternalServerError, err
	}

	if userID != nil {
		s.logger.Printf("Service CalculateTotalCost: success user_id=%s total_cost=%d", userID.String(), totalCost)
	} else {
		s.logger.Printf("Service CalculateTotalCost: success service_name=%s total_cost=%d", *req.ServiceName, totalCost)
	}
	return totalCost, http.StatusOK, nil
}
