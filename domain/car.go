package domain

import (
	"time"

	"pavel-vacha.cz/ctc/internal/paingas/types"
)

type Car struct {
	ID                int
	StationType       types.FuelType
	ArrivalAtStation  time.Time
	StationStartTime  time.Time
	StationEndTime    time.Time
	StationTime       time.Duration
	StationQueueTime  time.Duration
	ArrivalAtRegister time.Time
	RegisterStartTime time.Time
	RegisterEndTime   time.Time
	RegisterTime      time.Duration
	RegisterQueueTime time.Duration
}
