package main

import "fmt"

func main() {
	fmt.Println(ToCsvText([][]int{
		{0, 1, 2, 3, 45},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34}}))
}

func ToCsvText(array [][]int) string {
	var res string
	for i, a := range array {
		res += f(a)
		if i != len(array)-1 {
			res += "\n"
		}
	}
	return res
}
func f(arr []int) string {
	var str string
	for i, n := range arr {
		str += fmt.Sprintf("%d", n)
		if i != len(arr)-1 {
			str += ","
		}
	}
	return str
}
