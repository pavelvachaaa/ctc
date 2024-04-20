package utils

import (
	"math/rand"
	"time"
)

func GetRandomDuration(min, max int) time.Duration {
	return time.Duration(rand.Intn(max-min)+min) * time.Millisecond
}

func SimulateProcess(dur int) {
	time.Sleep(time.Duration(dur) * time.Millisecond)
}
