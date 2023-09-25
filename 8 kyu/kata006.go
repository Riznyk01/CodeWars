package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	_, err := fmt.Scan(&str)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Println(DNAtoRNA(str))
}

func DNAtoRNA(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}
