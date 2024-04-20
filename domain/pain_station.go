package domain

import "pavel-vacha.cz/ctc/internal/paingas/types"

type Station struct {
	ID           int
	Queue        chan *Car
	StationType  types.FuelType
	ServeTimeMin int
	ServeTimeMax int
}
