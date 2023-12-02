package main

import "fmt"

func main() {
	var dadYears, sonYears int
	_, err := fmt.Scan(&dadYears, &sonYears)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(TwiceAsOld(dadYears, sonYears))
}
func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	diff := dadYearsOld - 2*sonYearsOld
	if diff < 0 {
		return -diff
	}
	return diff
}
