package main

import "fmt"

func main() {
	var petalsNum int
	_, err := fmt.Scan(&petalsNum)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(HowMuchILoveYou(petalsNum))
}
func HowMuchILoveYou(i int) string {
	phrases := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}
	return phrases[(i-1)%6]
}
