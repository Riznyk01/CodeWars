package main

import "fmt"

func main() {
	fmt.Println(FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'}))
}
func FindMissingLetter(chars []rune) rune {
	var res rune
	for i := 0; i < len(chars); i++ {
		if chars[i] != chars[i+1]-1 {
			res = chars[i] + 1
			break
		}
	}
	return res
}
