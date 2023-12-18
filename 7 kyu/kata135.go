package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))
}

func SumOfIntegersInString(strng string) int {
	var num string
	sum := 0
	for i, r := range strng {
		if r >= '0' && r <= '9' {
			num += string(r)
			if i == len(strng)-1 {
				s(&num, &sum)
			}
		} else {
			s(&num, &sum)
		}
	}
	return sum
}
func s(num *string, sum *int) {
	e, _ := strconv.Atoi(*num)
	*sum += e
	*num = ""
}
