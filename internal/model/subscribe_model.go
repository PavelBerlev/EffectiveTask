package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	SubscribeModel struct {
		ID          int64
		ServiceName string
		Price       int
		UserID      uuid.UUID
		StartDate   time.Time
		EndDate     *time.Time
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
