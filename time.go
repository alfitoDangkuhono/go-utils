package timeutil

import (
	"fmt"
	"time"
)

func NowFormatted() string {
	return time.Now().Format(time.RFC1123)
}
func NowDateTimeWithLooping(format string) {
	for {
		fmt.Println("Time Right Now Start : " + time.Now().Format(format))
		time.Sleep(1 * time.Second)
	}
}
