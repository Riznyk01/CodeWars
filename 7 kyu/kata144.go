package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Vaporcode("Lets go to the movies"))
}

func Vaporcode(s string) string {
	return strings.Join(strings.Split(strings.ToUpper(strings.ReplaceAll(s, " ", "")), ""), "  ")
}
