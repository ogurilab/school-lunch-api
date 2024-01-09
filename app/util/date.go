package util

import "time"

func NowDate() string {
	currentTime := time.Now()

	return currentTime.Format("2006-01-02")
}

func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
