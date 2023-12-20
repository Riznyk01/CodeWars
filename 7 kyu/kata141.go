package main

import "fmt"

func main() {
	fmt.Println(OrderedCount("abracadabra"))
}

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	countMap := make(map[rune]int)
	for _, r := range text {
		countMap[r]++
	}
	arr := make([]Tuple, len(countMap))
	ind := 0
	for _, r := range text {
		if _, ok := countMap[r]; ok {
			arr[ind] = Tuple{r, countMap[r]}
			delete(countMap, r)
			ind++
		}
	}
	return arr
}
