package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(Disemvowel("This website is for losers LOL!"))
}
func Disemvowel(comment string) string {
	return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
}
