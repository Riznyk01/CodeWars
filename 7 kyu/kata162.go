package main

import "fmt"

func main() {
	fmt.Println(HasUniqueChar("  nAa"))
}
func HasUniqueChar(str string) bool {
	unics := make(map[rune]bool)
	for _, r := range str {
		if unics[r] {
			return false
		} else {
			unics[r] = true
		}
	}
	return true
}
