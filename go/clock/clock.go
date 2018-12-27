package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(hour int, minute int) Clock {
	clockHelper := Clock{hour, minute}
	return clockHelper.checkBounds()
}

func (clock Clock) Add(minute int) Clock {
	clock.minute += minute
	//	fmt.Printf("Clock minutes before bound check %d\n", clock.minute)
	return clock.checkBounds()
}

func (clock Clock) Subtract(minute int) Clock {
	// invert minutes
	minute = minute * -1
	return clock.Add(minute)
}

func (clock Clock) String() string {
	// in order for add to perform hours and minutes limit checks
	nc := clock.Add(0)
	hour := nc.hour
	minute := nc.minute
	var hs string
	var ms string
	if hour < 10 && hour >= 0 {
		hs = fmt.Sprintf("0%d", hour)
	} else {
		hs = fmt.Sprintf("%d", hour)
	}

	if minute < 10 && minute >= 0 {
		ms = fmt.Sprintf("0%d", minute)
	} else {
		ms = fmt.Sprintf("%d", minute)
	}

	return fmt.Sprintf("%s:%s", hs, ms)
}

func (clock Clock) checkBounds() Clock {
	// are we over 59 minutes
	if clock.minute > 59 {
		// next hour
		clock.hour = clock.hour + int(clock.minute/60)
		clock.minute = clock.minute % 60
	}
	// are we under 0 minutes
	if clock.minute < 0 {
		var sub = int(clock.minute / 60)
		if sub < 0 {
			sub = sub * -1
		}
		clock.hour = clock.hour - sub
		if (clock.minute % 60) != 0 {
			clock.hour = clock.hour - 1
		} else {
			clock.minute = clock.minute % 60

		}
		clock.minute = clock.minute % 60
		if clock.minute != 0 {
			clock.minute += 60
		}
	}

	// are we over 23 hours
	if clock.hour > 23 {
		clock.hour = clock.hour % 24
	}

	// are we under 0 hours
	if clock.hour < 0 {
		for clock.hour < 0 {
			clock.hour += 24
		}

		clock.hour = clock.hour % 24
	}

	return clock
}
