package dto

import "github.com/google/uuid"

type (
	SubscribeCreateRequest struct {
		ServiceName string    `json:"service_name" validate:"required"`
		Price       int       `json:"price" validate:"required,gt=0"`
		UserID      uuid.UUID `json:"user_id" validate:"required"`
		StartDate   string    `json:"start_date" validate:"required"`
		EndDate     *string   `json:"end_date,omitempty"`
	}

	SubscribeUpdateRequest struct {
		ID          int64     `json:"id" validate:"required"`
		ServiceName string    `json:"service_name" validate:"required"`
		Price       int       `json:"price" validate:"required"`
		UserID      uuid.UUID `json:"user_id" validate:"required"`
		StartDate   string    `json:"start_date" validate:"required"`
		EndDate     *string   `json:"end_date,omitempty"`
	}

	SubscribeTotalCostRequest struct {
		StartDate   string  `json:"start_date" validate:"required"`
		EndDate     string  `json:"end_date" validate:"required"`
		UserID      *string `json:"user_id"`
		ServiceName *string `json:"service_name"`
	}
)
