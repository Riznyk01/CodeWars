package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	name = strings.TrimSpace(name)
	fmt.Println(GreetForJohnny(name))
}

func GreetForJohnny(name string) string {
	if name == "Johnny" {
		return fmt.Sprint("Hello, my love!")
	}
	return fmt.Sprintf("Hello, %v!", name)
}
