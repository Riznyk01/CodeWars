package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Spacify("hello world"))
}
func Spacify(s string) string {
	return strings.Join(strings.Split(s, ""), " ")
}
