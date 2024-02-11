package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ScrabbleScore("street"))
}

func ScrabbleScore(st string) int {
	DictScores := map[string]int{}

	res := 0
	if len(st) == 0 {
		return 0
	}
	for _, r := range strings.Replace(st, " ", "", -1) {
		res += DictScores[strings.ToUpper(string(r))]
	}
	return res
}
