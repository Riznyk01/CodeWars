package main

import "fmt"

func main() {
	inp := []float32{5.0, 1.0, 1.0}
	fmt.Println(FindUniq(inp))
}

func FindUniq(arr []float32) float32 {
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	} else {
		for i := 1; i < len(arr); i++ {
			if arr[i] != arr[0] {
				return arr[i]
			}
		}
	}
	return 0
}
