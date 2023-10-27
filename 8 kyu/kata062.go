package main

import "fmt"

func main() {
	var empl, vac bool
	_, err := fmt.Scan(&empl, &vac)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(SetAlarm(empl, vac))
}
func SetAlarm(employed, vacation bool) bool {
	if employed && !vacation {
		return true
	}
	return false
}
