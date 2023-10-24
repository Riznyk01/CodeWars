package main

import (
	"fmt"
)

func main() {
	var width, height, depth int
	_, err := fmt.Scan(&width, &height, &depth)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(GetSize(width, height, depth))
}
func GetSize(w, h, d int) [2]int {
	area := 2 * (h*d + d*w + h*w)
	volume := w * h * d
	return [2]int{area, volume}
}
