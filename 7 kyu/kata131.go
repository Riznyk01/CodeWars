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
		fmt.Println(err)
	}
	fmt.Println(Accum(strings.TrimSpace(str)))
}
func Accum(s string) string {
	arr := make([]string, len([]rune(s)))
	for i, r := range []rune(s) {
		arr[i] = strings.ToUpper(string(r)) + strings.Repeat(strings.ToLower(string(r)), i)
	}
	return strings.Join(arr, "-")
}
