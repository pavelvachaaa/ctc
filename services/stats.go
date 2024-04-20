package services

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
	"pavel-vacha.cz/ctc/domain"
	"pavel-vacha.cz/ctc/internal/paingas/types"
)

// Funkce převezme čerpací stanici a vygeneruje yaml file dle zadání
// Vrací jednotlivé byty pro uložení do souborus
func GenerateStats(painGas *domain.PainGas) ([]byte, error) {
	fuelStats := make(map[types.FuelType]*types.FuelStats)
	cashRegisterStats := &types.RegisterStats{TotalCars: 0, TotalTime: 0, MaxQueueTime: 0}

	for _, car := range painGas.Cars {
		fuelTypeStats, err := fuelStats[car.StationType]
		if !err {
			fuelTypeStats = &types.FuelStats{}
			fuelStats[car.StationType] = fuelTypeStats
		}

		cashRegisterStats.TotalCars++
		fuelTypeStats.TotalCars++

		registerTime := car.RegisterTime + car.RegisterQueueTime
		cashRegisterStats.TotalTime += registerTime
		cashRegisterStats.MaxQueueTime = max(cashRegisterStats.MaxQueueTime, registerTime)

		serviceAndQueueTime := car.StationTime + car.StationQueueTime
		fuelTypeStats.TotalTime += serviceAndQueueTime
		fuelTypeStats.MaxQueueTime = max(fuelTypeStats.MaxQueueTime, serviceAndQueueTime)

	}

	// Output na std
	printerFunction(fuelStats, cashRegisterStats)

	result := struct {
		Stations  map[types.FuelType]*types.FuelStats `yaml:"stations"`
		Registers *types.RegisterStats                `yaml:"registers"`
	}{
		Stations: map[types.FuelType]*types.FuelStats{
			types.Gas:      fuelStats[types.Gas],
			types.Diesel:   fuelStats[types.Diesel],
			types.LPG:      fuelStats[types.LPG],
			types.Electric: fuelStats[types.Electric],
		},
		Registers: cashRegisterStats,
	}

	yamlBytes, err := yaml.Marshal(&result)
	if err != nil {
		return nil, err
	}

	return yamlBytes, nil
}

func printerFunction(fuelStats map[types.FuelType]*types.FuelStats, cashRegisterStats *types.RegisterStats) {
	for fuelType, stats := range fuelStats {
		if stats.TotalCars > 0 {
			avgTime := stats.TotalTime / time.Duration(stats.TotalCars)
			stats.AvgQueueTime = avgTime
			fmt.Printf("%s -> Počet aut: %d, Celkový čas: %s, Průměrný čas: %s, Max. strávený čas: %s\n",
				fuelType, stats.TotalCars, stats.TotalTime, avgTime, stats.MaxQueueTime)
		}
	}

	if cashRegisterStats.TotalCars > 0 {
		avgRegTime := cashRegisterStats.TotalTime / time.Duration(cashRegisterStats.TotalCars)
		cashRegisterStats.AvgQueueTime = avgRegTime
		fmt.Printf("\nPokladny:\nPočet aut: %d, Celkový čas: %s, Průměrný čas: %s, Max. strávený čas: %s\n",
			cashRegisterStats.TotalCars, cashRegisterStats.TotalTime, avgRegTime, cashRegisterStats.MaxQueueTime)
	}
}
