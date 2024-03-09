package main

import "fmt"

func main() {
	fmt.Println(Solution("abc"))
}

func Solution(str string) []string {
	s := str
	res := make([]string, (len(str)+1)/2)
	if len(s)%2 != 0 {
		s += "_"
	}
	for i := 0; i < len(s); i += 2 {
		res[i/2] = s[i : i+2]
	}
	return res
}
