package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(ReverseSeq(num))
}

func ReverseSeq(n int) []int {
	slice := make([]int, n)
	for i := 0; i < len(slice); i++ {
		slice[i] = n
		n--
	}
	return slice
}
