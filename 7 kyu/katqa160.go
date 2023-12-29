package main

import "fmt"

func main() {
	fmt.Println(Alternate(5, "true", "false"))
}
func Alternate(n int, firstValue string, secondValue string) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr[i] = firstValue
		} else {
			arr[i] = secondValue
		}
	}
	return arr
}
