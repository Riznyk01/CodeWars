package main

import "fmt"

func main() {
	var numbers [2]int
	_, err1 := fmt.Scan(&numbers[0])
	if err1 != nil {
		fmt.Println(err1)
	}
	_, err2 := fmt.Scan(&numbers[1])
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(ReduceFraction(numbers))
}
func ReduceFraction(fraction [2]int) [2]int {
	GCD := 0
	n := fraction[0]
	d := fraction[1]
	if n < d {
		n, d = d, n
	}
	for GCD == 0 {
		switch {
		case n%d == 0:
			GCD = d
			break
		case n%d != 0:
			n, d = d, n%d
		}
	}
	return [2]int{fraction[0] / GCD, fraction[1] / GCD}
}
