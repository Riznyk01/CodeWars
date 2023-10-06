package main

import (
	"fmt"
	"time"
)

func main() {
	var sec int
	_, err := fmt.Scan(&sec)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(HumanReadableTime(sec))
}
func HumanReadableTime(seconds int) string {
	d := time.Second * time.Duration(seconds)
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	timeString := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	return timeString
}
