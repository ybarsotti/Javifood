package pkg

import "time"

const (
    TimeFormat = "2000-01-02T00:00:00.000Z08:00"
)

func ParseDate(timeStr string) (time.Time, error) {
    return time.Parse(TimeFormat, timeStr)
}
