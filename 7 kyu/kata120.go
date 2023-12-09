package main

import "fmt"

func main() {
	s := []string{"beg", "life", "i", "to"}
	fmt.Println(SortByLength(s))
}
func SortByLength(arr []string) []string {
	res := make([]string, len(arr))
	for len(arr) > 0 {
		idx, maxLen := 0, 0
		for i, w := range arr {
			if len(w) > maxLen {
				maxLen = len(w)
				idx = i
			}
		}
		res[len(arr)-1] = arr[idx]
		arr[idx] = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
	}
	return res
}
