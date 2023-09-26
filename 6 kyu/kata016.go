package main

import (
	"fmt"
	"strings"
)

func main() {
	inpArr := []string{"abracadabra", "allottee", "assessee"}
	fmt.Println(Dup(inpArr))
}

func Dup(arr []string) []string {
	var outArr []string
	var bf rune

	filter := func(r rune) rune {
		if r != bf {
			bf = r
			return r
		}
		return -1
	}

	for _, input := range arr {
		bf = 0
		outArr = append(outArr, strings.Map(filter, input))
	}
	return outArr
}
