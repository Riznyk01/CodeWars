package main

import "fmt"

func main() {
	var xOne, xTwo int
	_, err := fmt.Scan(&xOne, &xTwo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Quadratic(xOne, xTwo))
}

func Quadratic(x1, x2 int) [3]int {
	return [3]int{1, -x1 - x2, x1 * x2}
}
