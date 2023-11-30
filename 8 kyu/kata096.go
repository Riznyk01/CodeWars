package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	array := []any{9, 3, "7", "3"}
	fmt.Println(SumMix(array))
}

func SumMix(arr []any) int {
	var sum int
	for _, n := range arr {
		switch v := n.(type) {
		case int:
			sum += v
		case string:
			vInt, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			sum += vInt
		}

	}
	return sum
}
