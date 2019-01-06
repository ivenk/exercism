package clock

import "fmt"

// Clock respresents a real world clock, but misses seconds
type Clock struct {
	minute int
}

// New initialises a new clock with the given time
func New(hour int, minute int) Clock {
	var min = hour * 60 // may be negativ
	min += minute       // again may be negativ
	clockHelper := Clock{min}
	return clockHelper.checkBounds()
}

// Add adds the given amount of minutes to the clock
func (clock Clock) Add(minute int) Clock {
	clock.minute += minute
	return clock.checkBounds()
}

// Subtract subtracts the given amount of minutes from the clock
func (clock Clock) Subtract(minute int) Clock {
	// invert minutes
	minute = minute * -1
	return clock.Add(minute)
}

func (clock Clock) String() string {
	hour := clock.minute / 60
	min := clock.minute % 60
	return fmt.Sprintf("%02d:%02d", hour, min)
}

// checkBounds enforces the time format
func (clock Clock) checkBounds() Clock {
	minPerDay := 24 * 60

	if clock.minute < 0 {
		var days int
		days = (-1 * (clock.minute / minPerDay)) + 1
		clock.minute += minPerDay * days
	}

	for clock.minute >= minPerDay {
		clock.minute -= minPerDay
	}

	return clock
}
