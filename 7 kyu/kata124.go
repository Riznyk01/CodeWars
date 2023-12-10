package main

import "fmt"

func main() {
	var a, num int
	_, err := fmt.Scan(&a, &num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(OverTheRoad(a, num))
}

func OverTheRoad(address int, n int) int {
	return 2*(n+1-(address+1)/2) - (address+1)%2
}
