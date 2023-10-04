package main

import "fmt"

func main() {
	var r, g, b int
	_, err := fmt.Scan(&r, &g, &b)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(RGB(r, g, b))
}

func RGB(r, g, b int) string {
	fc := func(n int) int {
		if n < 0 {
			n = 0
		} else if n > 255 {
			n = 255
		}
		return n
	}
	return fmt.Sprintf("%02X%02X%02X", fc(r), fc(g), fc(b))
}
