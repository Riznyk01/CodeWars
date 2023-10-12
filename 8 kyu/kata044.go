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
	fmt.Println(ReverseWords(str))
}

func ReverseWords(str string) string {
	wordsArr := strings.Split(str, " ")
	wordsRevArr := make([]string, len(wordsArr))
	for i := len(wordsArr) - 1; i >= 0; i-- {
		wordsRevArr[len(wordsArr)-1-i] = wordsArr[i]
	}
	return strings.Join(wordsRevArr, " ")
}
