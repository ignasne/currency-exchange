package datasource

import "time"

const timeFormat = "2006-01-02 15:04:05"

func Timestamp() string {
	return time.Now().UTC().Format(timeFormat)
}

func TimeToString(time time.Time) string {
	return time.Format(timeFormat)
}
