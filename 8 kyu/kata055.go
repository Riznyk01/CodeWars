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
	fmt.Println(AbbrevName(str))
}
func AbbrevName(name string) string {
	nameSlice := strings.Split(name, " ")
	for i := 0; i < len(nameSlice); i++ {
		if nameSlice[i][0] >= 'a' && nameSlice[i][0] <= 'z' {
			nameSlice[i] = strings.ToUpper(nameSlice[i])
		}
	}
	return fmt.Sprintf("%c.%c", nameSlice[0][0], nameSlice[1][0])
}
