package util

import "time"

const LayoutDate = "2006-01-02T15:04:05"

func StringToTime(stringDate, layoutDate string) (t time.Time, err error) {
	return time.Parse(layoutDate, stringDate)
}
