package clock

import (
	"fmt"
	"math"
)

func GenerateTimeInterval(x int) (times []string) {
	tt := 0

	for i := 0; tt < 24*60; i++ {
		hh := math.Floor(float64(tt / 60)) // getting hours of day in 0-24 format
		mm := tt % 60                      // getting minutes of the hour in 0-55 format

		hourFormat := fmt.Sprintf("%02d", int(hh))
		minuteFormat := fmt.Sprintf("%02d", mm)

		if minuteFormat == "0" {
			minuteFormat = "00"
		}

		times = append(times, fmt.Sprintf("%v:%v", hourFormat, minuteFormat))
		tt = tt + x
	}

	return
}
