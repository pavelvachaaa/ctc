package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"pavel-vacha.cz/ctc/domain"
	"pavel-vacha.cz/ctc/internal/paingas/types"
	"pavel-vacha.cz/ctc/internal/paingas/utils"
	"pavel-vacha.cz/ctc/services"
)

func initPainGas(config types.Configuration) *domain.PainGas {
	stations, stationWorkGroups := services.InitStations(config)
	registers, registerWorkGroups := services.InitCashRegisters(config)

	return &domain.PainGas{
		Stations:           stations,
		StationWorkGroups:  stationWorkGroups,
		RegisterWorkGroups: *registerWorkGroups,
		Registers:          registers,
	}
}

func main() {
	config, err := utils.GetConfig("./configs/config.dev.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	painGas := initPainGas(*config)
	painGas.Cars = []*domain.Car{}

	for _, station := range painGas.Stations {
		go services.StationRoutine(station, painGas)
	}

	for _, register := range painGas.Registers {
		go services.CashRegisterRoutine(register, painGas)
	}

	services.SimulateCars(painGas, *config)

	painGas.CarsWorkGroup.Wait()

	defer func() {
		for _, station := range painGas.Stations {
			close(station.Queue)
		}

		for _, register := range painGas.Registers {
			close(register.Queue)
		}
	}()

	yamlBytes, err := services.GenerateStats(painGas)
	if err != nil {
		fmt.Println("Error generating YAML:", err)
		return
	}

	outputDir := "output"
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	filePath := filepath.Join(outputDir, "stats.yaml")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(yamlBytes)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}
