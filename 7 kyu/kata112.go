package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(DNAStrand(str))
}
func DNAStrand(dna string) string {
	r := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
	return r.Replace(dna)
}
