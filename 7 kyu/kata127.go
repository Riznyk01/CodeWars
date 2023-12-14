package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ToJadenCase(str))
}
func ToJadenCase(str string) string {
	resArr := make([]string, len(strings.Split(str, " ")))
	for i, w := range strings.Split(str, " ") {
		runes := []rune(w)
		resArr[i] = string(unicode.ToUpper(runes[0])) + string(runes[1:])
	}
	return strings.Join(resArr, " ")
}
