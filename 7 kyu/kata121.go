package main

import "fmt"

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(GetMiddle(str))
}
func GetMiddle(s string) string {
	n := len(s)
	return s[(n/2 + (n%2 - 1)) : (n/2)+1]
}
