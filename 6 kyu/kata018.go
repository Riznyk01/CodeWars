package main

import (
	"fmt"
	"strings"
)

func main() {
	var inpStr string
	_, err := fmt.Scan(&inpStr)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(High(inpStr))
}

type WordSumPair struct {
	Word string
	Sum  int
}

func High(s string) string {
	words := strings.Fields(s)
	var wordSumPairs []WordSumPair
	var maxWord string
	maxSum := -1

	for _, word := range words {
		var sum int
		for _, r := range word {
			sum += int(r) - 96
		}
		pair := WordSumPair{Word: word, Sum: sum}
		wordSumPairs = append(wordSumPairs, pair)
	}
	for _, pair := range wordSumPairs {
		if pair.Sum > maxSum {
			maxSum = pair.Sum
			maxWord = pair.Word
		} else if pair.Sum == maxSum {
			continue
		}
	}
	return maxWord
}
