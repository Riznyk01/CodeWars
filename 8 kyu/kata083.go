package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter values in the format \"1:0\",\"2:0\",\"3:0\",...:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	gamesSlice := strings.Split(input, ",")
	fmt.Println(Points(gamesSlice))
}
func Points(games []string) int {
	p := 0
	for _, score := range games {
		pair := strings.Split(score, ":")
		if pair[0] > pair[1] {
			p += 3
		} else if pair[0] > pair[1] {
			continue
		} else if pair[0] == pair[1] {
			p += 1
		}
	}
	return p
}
