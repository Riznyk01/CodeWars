package main

import "fmt"

func main() {
	var num uint64
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Digits(num))
}
func Digits(n uint64) int {
	cnt := 1
	for n > 9 {
		n /= 10
		cnt++
	}
	return cnt
}
