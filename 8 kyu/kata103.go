package main

import "fmt"

func main() {
	var aA, bB bool
	_, err := fmt.Scan(&aA, &bB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Xor(aA, bB))
}
func Xor(a, b bool) bool {
	return a != b
}
