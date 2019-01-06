package clock

import "fmt"

// Clock respresents a real world clock, but misses seconds
type Clock int

// New initialises a new clock with the given time
func New(hour int, minute int) Clock {
	var min = (hour * 60) + minute // may be negativ
	minPerDay := 24 * 60

	if min < 0 {
		var days int
		days = (-1 * (min / minPerDay)) + 1
		min += minPerDay * days
	}

	for min >= minPerDay {
		min -= minPerDay
	}

	return Clock(min)
}

// Add adds the given amount of minutes to the clock
func (clock Clock) Add(minute int) Clock {
	return New(0, int(clock)+minute)
}

// Subtract subtracts the given amount of minutes from the clock
func (clock Clock) Subtract(minute int) Clock {
	// invert minutes
	minute = minute * -1
	return New(0, int(clock)+minute)
}

func (clock Clock) String() string {
	hour := int(clock) / 60
	min := int(clock) % 60
	return fmt.Sprintf("%02d:%02d", hour, min)
}
