package main

import "fmt"

func main() {
	nums := []bool{true, false, true}
	fmt.Println(CountSheeps(nums))
}
func CountSheeps(numbers []bool) int {
	cnt := 0
	for i := range numbers {
		if numbers[i] == true {
			cnt++
		}
	}
	return cnt
}
