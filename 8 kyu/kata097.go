package main

import (
	"fmt"
	"sort"
)

func main() {
	array1 := []int{1, 3, 5, 7, 9}
	array2 := []int{10, 8, 6, 4, 2}
	fmt.Println(MergeArrays(array1, array2))
}

func MergeArrays(arr1, arr2 []int) []int {
	sumArr := make([]int, len(arr1)+len(arr2))
	sumArr = append(arr1, arr2...)
	resultArr := make([]int, 0)
	sort.Ints(sumArr)
	resultArr = append(resultArr, sumArr[0])
	for i := 1; i < len(sumArr); i++ {
		if sumArr[i] != sumArr[i-1] {
			resultArr = append(resultArr, sumArr[i])
		}
	}
	return resultArr
}
