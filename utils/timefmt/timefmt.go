package timefmt

import "time"

const localTimezoneLocation = "Asia/Jakarta"

func TimeInLocalTimezone(t time.Time) time.Time {
	loc, err := time.LoadLocation(localTimezoneLocation)
	if err != nil {
		panic(err)
	}
	return t.In(loc)
}
