package timeutil

import (
	"fmt"
	"time"
)

func NowFormatted() string {
	return time.Now().Format(time.RFC1123)
}
func NowDateTimeWithLooping(format string) string {
	return "Time Right Now Start : " + time.Now().Format(format)
}
