package dates

import "time"

const dateLayout = "2006-01-02T15:04:05"

func GetUTCTime() time.Time {
	return time.Now().UTC()
}

func GetUTCString() string {
	return GetUTCTime().Format(dateLayout)
}
