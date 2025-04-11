package helperz

import (
	"os"
	"time"
)

type DatetimeHelper interface {
	GetNow(duration time.Duration) *time.Time
	ParseStringToTime(format, dateStr string, loc *time.Location) (time.Time, error)
	ParseStringToTimeWithTimezone(format, dateStr string) (time.Time, error)
	TimeToString(t *time.Time) string
}

func GetNow(duration time.Duration) *time.Time {
	now := time.Now().Add(duration)
	return &now
}

func ParseStringToTime(format, dateStr string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		return time.Parse(format, dateStr)
	}
	return time.ParseInLocation(format, dateStr, loc)
}

func ParseStringToTimeWithTimezone(format, dateStr string) (time.Time, error) {
	timezone := os.Getenv("TZ")
	if timezone == "" {
		timezone = "Asia/Jakarta"
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return ParseStringToTime(format, dateStr, loc)
}

func TimeToString(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
