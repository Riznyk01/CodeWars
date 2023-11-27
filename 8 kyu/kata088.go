package main

import (
	"fmt"
	"time"
)

func main() {
	var hours, mins, secs int
	_, err := fmt.Scan(&hours, &mins, &secs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Past(hours, mins, secs))
}
func Past(h, m, s int) int {
	duration := time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second
	return int(duration.Milliseconds())
}
