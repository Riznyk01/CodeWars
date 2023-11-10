package main

import "fmt"

func main() {
	var sal int
	var bon bool
	_, err := fmt.Scan(&sal, &bon)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(BonusTime(sal, bon))
}

func BonusTime(salary int, bonus bool) string {
	if bonus {
		salary *= 10
	}
	return fmt.Sprintf("£%d", salary)
}
