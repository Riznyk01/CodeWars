package main

import "fmt"

func main() {
	grains := make(chan string, 7)
	grains <- "good"
	grains <- "bad"
	grains <- "bad"
	grains <- "good"
	grains <- "bad"
	grains <- "bad"
	grains <- "good"
	close(grains)
	g, b := PickGrains(grains)
	fmt.Printf("good %v, bad %v", g, b)
}
func PickGrains(grains <-chan string) (good int, bad int) {
	messages := make(map[string]int)
	for v := range grains {
		messages[v]++
	}
	return messages["good"], messages["bad"]
}
