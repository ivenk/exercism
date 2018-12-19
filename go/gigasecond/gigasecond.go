package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	// we add one gigasecond to the current time
	return t.Add(time.Second * 1000000000)

}
