package types

import "time"

type FuelType string

const (
	Diesel   FuelType = "diesel"
	Electric FuelType = "electric"
	Gas      FuelType = "gas"
	LPG      FuelType = "lpg"
)

type StationConfiguration struct {
	Type         string
	Count        int
	ServeTimeMin DurationConfiguration `yaml:"serve_time_min"`
	ServeTimeMax DurationConfiguration `yaml:"serve_time_max"`
}

type DurationConfiguration struct {
	Duration time.Duration
}

// Přetížení metody pro parsování na Duration pro zjednodušení práce s configem.
func (d *DurationConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw string
	if err := unmarshal(&raw); err != nil {
		return err
	}
	duration, err := time.ParseDuration(raw)
	if err != nil {
		return err
	}
	d.Duration = duration
	return nil
}

type Configuration struct {
	Cars struct {
		Count          int
		ArrivalTimeMin DurationConfiguration `yaml:"arrival_time_min"`
		ArrivalTimeMax DurationConfiguration `yaml:"arrival_time_max"`
	} `yaml:"cars"`

	Stations map[FuelType]StationConfiguration

	Registers struct {
		Count         int
		HandleTimeMin DurationConfiguration `yaml:"handle_time_min"`
		HandleTimeMax DurationConfiguration `yaml:"handle_time_max"`
	} `yaml:"registers"`
}

type FuelStats struct {
	TotalCars    int           `yaml:"total_cars"`
	TotalTime    time.Duration `yaml:"total_time"`
	MaxQueueTime time.Duration `yaml:"max_queue_time"`
	AvgQueueTime time.Duration `yaml:"avg_queue_time"`
}

type RegisterStats struct {
	TotalCars    int           `yaml:"total_cars"`
	TotalTime    time.Duration `yaml:"total_time"`
	MaxQueueTime time.Duration `yaml:"max_queue_time"`
	AvgQueueTime time.Duration `yaml:"avg_queue_time"`
}
