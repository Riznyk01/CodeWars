package main

import "fmt"

func main() {
	var pos, rol int

	_, err := fmt.Scan(&pos, &rol)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Move(pos, rol))
}

func Move(position int, roll int) int {
	return position + roll*2
}
