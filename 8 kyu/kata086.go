package main

import (
	"fmt"
)

func main() {

	var num int

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PowersOfTwo(num))
}

func PowersOfTwo(n int) []uint64 {
	nums := make([]uint64, n+1)
	nums[0] = 1
	for i := 1; i <= n; i++ {
		//nums[i] = uint64(math.Pow(2, float64(i)))
		nums[i] = nums[i-1] * 2
	}
	return nums
}
