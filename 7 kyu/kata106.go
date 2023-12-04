package main

import "fmt"

func main() {
	var y int
	_, err := fmt.Scan(&y)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(IsLeapYear(y))
}
func IsLeapYear(year int) bool {
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		return true
	}
	return false
}
