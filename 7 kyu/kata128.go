package main

import "fmt"

func main() {
	numsArr := []int{1, -1, 2, -2, 3}
	fmt.Println(SolveElemParity(numsArr))
}
func SolveElemParity(arr []int) (res int) {
	numsMap := make(map[int]struct{})
	for _, n := range arr {
		_, exist := numsMap[-n]
		if exist {
			delete(numsMap, -n)
		} else {
			numsMap[n] = struct{}{}
		}
	}
	for k := range numsMap {
		res = k
	}
	return res
}
