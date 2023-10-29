package main

import "fmt"

func main() {
	var llGoals, cdrGoals, clGoals int
	_, err := fmt.Scan(&llGoals, &cdrGoals, &clGoals)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(Goals(llGoals, cdrGoals, clGoals))
}
func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}
