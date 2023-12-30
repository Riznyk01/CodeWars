package main

import "fmt"

func main() {
	fmt.Println(HasUniqueChar("  nAa"))
}
func HasUniqueChar(str string) bool {
	unics := make(map[rune]struct{})
	for _, r := range str {
		if _, ex := unics[r]; ex {
			return false
		} else {
			unics[r] = struct{}{}
		}
	}
	return true
}
