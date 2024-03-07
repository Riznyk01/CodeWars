package main

import "fmt"

func main() {
	fmt.Println(CleanString("abc#d##c"))
}
func CleanString(s string) string {
	out := ""
	for _, r := range s {
		if r == '#' {
			if len(out) > 0 {
				out = out[:len(out)-1]
			}
		} else {
			out += string(r)
		}
	}
	return out
}
