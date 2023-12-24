package main

import "fmt"

func main() {
	fmt.Println(bandNameGenerator("alaska"))
}
func bandNameGenerator(word string) string {
	wordArr := []rune(word)
	wordArr[0] -= 32
	if wordArr[0]+32 == wordArr[len(wordArr)-1] {
		wordArr = append(wordArr, wordArr[1:]...)
	} else {
		for i, r := range wordArr {
			if r == '-' {
				wordArr[i+1] -= 32
			}
		}
		return "The " + string(wordArr)
	}
	return string(wordArr)
}
