package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(NbDig(550, 5))
}
func NbDig(n int, d int) (cnt int) {
	for i := 0; i <= n; i++ {
		cnt += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
	return cnt
}
