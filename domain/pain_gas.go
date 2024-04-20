package domain

import (
	"sync"

	"pavel-vacha.cz/ctc/internal/paingas/types"
)

type PainGas struct {
	Stations           []*Station
	Registers          []*Register
	CarsWaitGroup      sync.WaitGroup
	StationWorkGroups  map[types.FuelType]*sync.WaitGroup
	RegisterWorkGroups sync.WaitGroup
	Cars               []*Car
}
