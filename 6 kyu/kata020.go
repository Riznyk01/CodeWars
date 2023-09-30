package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(Persistence(num))
}
func Persistence(n int) int {
	cnt := 0

	fc := func(n int) int {
		m := 1
		for i := 1; n/i > 0; i *= 10 {
			m = m * ((n / i) % 10)
		}
		return m
	}

	if n >= 10 {
		for n > 9 {
			n = fc(n)
			cnt++
		}
	}
	return cnt
}
