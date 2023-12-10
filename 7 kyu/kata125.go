package main

import (
	"fmt"
	"time"
)

func main() {
	var y int
	_, err := fmt.Scan(&y)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UnluckyDays(y))
}
func UnluckyDays(year int) (cnt int) {
	for m := 1; m <= 12; m++ {
		if time.Date(year, time.Month(m), 13, 0, 0, 0, 0, time.UTC).Weekday() == time.Friday {
			cnt++
		}
	}
	return cnt
}
