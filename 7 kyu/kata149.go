package main

import "fmt"

func main() {
	fmt.Println(AreCoprime(12, 39))
}

func AreCoprime(n, m int) bool {
	if m < n {
		n, m = m, n
	}
	for i := 0; n-i > 1; i++ {
		if m%(n-i) == 0 && n%(n-i) == 0 {
			return false
		}
	}
	return true
}
