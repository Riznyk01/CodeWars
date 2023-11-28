package main

import "fmt"

func main() {
	var bull, drag int

	_, err := fmt.Scan(&bull, &drag)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Hero(bull, drag))
}
func Hero(bullets, dragons int) bool {
	if bullets/dragons >= 2 {
		return true
	}
	return false
}
