package main

import "fmt"

func main() {
	fmt.Println(LateRide(240))
}
func LateRide(n int) int {
	h := n / 60
	m := n - h*60
	return h/10 + h%10 + m/10 + m%10
}
