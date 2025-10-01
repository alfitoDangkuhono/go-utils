package timeutil

import (
	"fmt"
	"time"
)

func NowFormatted() string {
	return time.Now().Format(time.RFC1123)
}
func NowDateTimeWithLooping(format string) string {
	// for {
		
	// 	time.Sleep(1 * time.Second)
	// }
	fmt.Println("Time Right Now Start : " + time.Now().Format(format))
}
