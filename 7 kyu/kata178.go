package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RemoveDuplicateWords("alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta"))
}
func RemoveDuplicateWords(str string) string {
	arr := make([]string, 0)
	unic := make(map[string]struct{})
	ind := 0
	for _, w := range strings.Split(strings.TrimSpace(str), " ") {
		if _, ex := unic[w]; !ex {
			unic[w] = struct{}{}
			arr = append(arr, w)
			ind++
		}
	}
	return strings.Join(arr, " ")
}
