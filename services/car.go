package services

import (
	"math/rand"
	"time"

	"pavel-vacha.cz/ctc/domain"
	"pavel-vacha.cz/ctc/internal/paingas/types"
	"pavel-vacha.cz/ctc/internal/paingas/utils"
)

func SimulateCars(painGas *domain.PainGas, config types.Configuration) {
	for i := 0; i < config.Cars.Count; i++ {
		car := &domain.Car{
			ID:               i,
			StationType:      utils.GetRandomFuelType(),
			ArrivalAtStation: time.Now(),
		}

		painGas.Cars = append(painGas.Cars, car)
		painGas.CarsWorkGroup.Add(1)

		station := GetStationWithShortestQueue(painGas.Stations, car.StationType)
		station.Queue <- car

		arrivalTime := rand.Intn(int(config.Cars.ArrivalTimeMax.Duration.Milliseconds())-int(config.Cars.ArrivalTimeMin.Duration.Milliseconds())+1) + int(config.Cars.ArrivalTimeMin.Duration.Milliseconds())
		utils.SimulateProcess(arrivalTime)
	}
}
