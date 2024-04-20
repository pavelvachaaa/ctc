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

func InitCashRegisters(config types.Configuration) ([]*domain.Register, *sync.WaitGroup) {
	var registers []*domain.Register
	var registerWorkGroups sync.WaitGroup
	for i := 0; i < config.Registers.Count; i++ {
		registerWorkGroups.Add(1)
		registers = append(registers, &domain.Register{
			ID:            i,
			Queue:         make(chan *domain.Car, 20),
			HandleTimeMin: int(config.Registers.HandleTimeMin.Duration.Milliseconds()),
			HandleTimeMax: int(config.Registers.HandleTimeMax.Duration.Milliseconds()),
		})
	}

	return registers, &registerWorkGroups
}

func CashRegisterRoutine(register *domain.Register, gs *domain.PainGas) {
	for car := range register.Queue {
		car.RegisterStartTime = time.Now()

		handleTime := utils.GetRandomDuration(register.HandleTimeMin, register.HandleTimeMax)
		car.RegisterTime = handleTime
		time.Sleep(handleTime)

		car.RegisterEndTime = time.Now()
		car.RegisterQueueTime = time.Since(car.RegisterStartTime) - handleTime

		gs.CarsWaitGroup.Done()
	}
}

func GetRegisterWithShortestQueue(registers []*domain.Register) *domain.Register {
	var shortestQueue *domain.Register
	shortestQueueLength := math.MaxInt32 // Initialize with maximum possible value

	for _, register := range registers {
		queueLength := len(register.Queue)
		if queueLength < shortestQueueLength {
			shortestQueue = register
			shortestQueueLength = queueLength
		} else if queueLength == shortestQueueLength && rand.Intn(2) == 0 {
			shortestQueue = register
		}
	}

	return shortestQueue
}
