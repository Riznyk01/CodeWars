package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Calc("abcdef"))
}
func Calc(s string) int {
	var total1Str, total2Str string
	for _, r := range s {
		total1Str += fmt.Sprintf("%v", r)
	}
	total2Str = strings.ReplaceAll(total1Str, "7", "1")
	return SumStr(total1Str) - SumStr(total2Str)
}
func SumStr(str string) (res int) {
	for _, r := range str {
		n, _ := strconv.Atoi(string(r))
		res += n
	}
	return res
}
