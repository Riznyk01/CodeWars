package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(HowManyDalmatians(num))
}

func HowManyDalmatians(number int) string {
	if number <= 10 {
		return "Hardly any"
	} else if number <= 50 {
		return "More than a handful!"
	} else if number == 101 {
		return "101 DALMATIONS!!!"
	}
	return "Woah that's a lot of dogs!"
}
