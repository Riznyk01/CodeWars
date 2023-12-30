package main

import "fmt"

func main() {
	fmt.Println(inviteMoreWomen([]int{1, -1, 1}))
}
func inviteMoreWomen(L []int) bool {
	sum := 0
	for _, n := range L {
		sum += n
	}
	return sum > 0
}
