package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	phrase, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(Solution(phrase))
}
func Solution(word string) string {
	bytesWord := []byte(word)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		bytesWord[i], bytesWord[j] = bytesWord[j], bytesWord[i]
	}
	return string(bytesWord)
}
