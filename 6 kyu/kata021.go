package main

import (
	"fmt"
	"strings"
)

func main() {

	var num int
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	for _, value := range TowerBuilder(num) {
		fmt.Println(value)
	}
}
func TowerBuilder(nFloors int) []string {
	tower := make([]string, nFloors)
	for i := 0; i < nFloors; i++ {
		tower[i] = strings.Repeat(" ", nFloors-i-1) + strings.Repeat("*", 2*i+1) + strings.Repeat(" ", nFloors-i-1)
	}
	return tower
}
