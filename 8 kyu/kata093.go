package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
	fmt.Println(TwoSort(words))
}

func TwoSort(arr []string) string {
	sort.Strings(arr)
	l := strings.Split(arr[0], "")
	return strings.Join(l, "***")
}
