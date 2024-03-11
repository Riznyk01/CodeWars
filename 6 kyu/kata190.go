package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func CreatePhoneNumber(numbers [10]uint) (r string) {
	for _, n := range numbers {
		r += strconv.Itoa(int(n))
	}
	return fmt.Sprintf("(%s) %s-%s", r[0:3], r[3:6], r[6:10])
}
