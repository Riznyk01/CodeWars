package main

import (
	"fmt"
)

func main() {
	var pOne, pTwo string
	_, err := fmt.Scan(&pOne, &pTwo)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(Rps(pOne, pTwo))
}
func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}
	if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	}
	if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	}
	if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	}
	return "Player 2 won!"
}
