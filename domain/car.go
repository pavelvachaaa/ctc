package domain

import (
	"time"

	"pavel-vacha.cz/ctc/internal/paingas/types"
)

type Car struct {
	ID                int
	StationType       types.FuelType
	ArrivalAtStation  time.Time
	ServiceStartTime  time.Time
	ServiceEndTime    time.Time
	ServiceTime       time.Duration
	ServiceQueueTime  time.Duration
	ArrivalAtReg      time.Time
	RegisterStartTime time.Time
	RegisterEndTime   time.Time
	RegisterTime      time.Duration
	RegisterQueueTime time.Duration
}
