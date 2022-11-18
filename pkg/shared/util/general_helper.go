package util

import "time"

func GetCurrentDate() time.Time {
	return time.Now().Local().Add(time.Hour * time.Duration(7))
}
