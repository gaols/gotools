package gotools

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// TimeUnitTypeSecond as second
	TimeUnitTypeSecond int = 0
	// TimeUnitTypeMinute as minute
	TimeUnitTypeMinute int = 1
	// TimeUnitTypeHour as hour
	TimeUnitTypeHour int = 2
)

// RandDuration generate pseudo time duration between lower and upper with lower and upper bound inclusively.
func RandDuration(lower, upper, timeUnitType int) time.Duration {
	if upper < lower {
		panic("upper bound should gt lower bound")
	}

	d := time.Second
	switch timeUnitType {
	case TimeUnitTypeSecond:
		d = time.Second
	case TimeUnitTypeMinute:
		d = time.Minute
	case TimeUnitTypeHour:
		d = time.Hour
	default:
		panic(fmt.Sprintf("invalid time unit type: %d", timeUnitType))
	}

	return time.Duration(rand.Intn(upper-lower)+lower) * d
}
