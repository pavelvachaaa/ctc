package services

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"pavel-vacha.cz/ctc/domain"
	"pavel-vacha.cz/ctc/internal/paingas/types"
	"pavel-vacha.cz/ctc/internal/paingas/utils"
)

func InitStations(config types.Configuration) ([]*domain.Station, map[types.FuelType]*sync.WaitGroup) {
	var stations []*domain.Station
	stationWorkGroups := make(map[types.FuelType]*sync.WaitGroup)

	for fuelType, station := range config.Stations {
		stationWorkGroups[fuelType] = &sync.WaitGroup{}
		for i := 0; i < station.Count; i++ {
			// Indikujeme, že jsme přidali novou stanici
			stationWorkGroups[fuelType].Add(1)
			stations = append(stations, &domain.Station{
				ID:           i,
				Queue:        make(chan *domain.Car, 5), // Buffered kanál pro auta
				StationType:  fuelType,
				ServeTimeMin: int(station.ServeTimeMin.Duration.Milliseconds()),
				ServeTimeMax: int(station.ServeTimeMax.Duration.Milliseconds()),
			})
		}
	}
	return stations, stationWorkGroups
}

func GetStationWithShortestQueue(stations []*domain.Station, carType types.FuelType) *domain.Station {
	var shortestQueue *domain.Station
	shortestQueueLength := math.MaxInt32

	for _, station := range stations {
		if station.StationType == carType {
			queueLength := len(station.Queue)
			if queueLength < shortestQueueLength {
				shortestQueue = station
				shortestQueueLength = queueLength
			} else if queueLength == shortestQueueLength && rand.Intn(2) == 0 {
				shortestQueue = station
			}
		}
	}

	return shortestQueue
}

func StationRoutine(station *domain.Station, painGas *domain.PainGas) {
	for car := range station.Queue {

		car.ServiceStartTime = time.Now()
		serveTime := utils.GetRandomDuration(station.ServeTimeMin, station.ServeTimeMax)
		car.ServiceTime = serveTime

		time.Sleep(serveTime)

		car.ServiceEndTime = time.Now()
		car.ServiceQueueTime = time.Since(car.ServiceStartTime) - serveTime

		// Pošli auto do ke kase, kde je nejmenší fronta
		register := GetRegisterWithShortestQueue(painGas.Registers)
		car.ArrivalAtReg = time.Now()
		register.Queue <- car
	}
}
