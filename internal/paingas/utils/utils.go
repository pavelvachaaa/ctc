package utils

import (
	"math/rand"

	"pavel-vacha.cz/ctc/internal/paingas/types"
)

func GetRandomFuelType() types.FuelType {
	fuelTypes := []types.FuelType{types.Gas, types.Diesel, types.LPG, types.Electric}
	return fuelTypes[rand.Intn(len(fuelTypes))]
}
