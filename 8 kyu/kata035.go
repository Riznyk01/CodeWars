package main

import (
	"fmt"
)

func main() {
	var l, w, h float64
	_, err := fmt.Scan(&l, &w, &h)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(GetVolumeOfCuboid(l, w, h))
}
func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}
