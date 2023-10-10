package main

import (
	"fmt"
)

func main() {
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(Greet(name))
}
func Greet(name string) string {
	return fmt.Sprintf("Hello, " + name + " how are you doing today?")
}
