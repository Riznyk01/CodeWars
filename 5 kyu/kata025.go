package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var startingAddresses, finalAddresses string
	_, err := fmt.Scan(&startingAddresses, &finalAddresses)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(IpsBetween(startingAddresses, finalAddresses))
}

func IpsBetween(start, end string) int {
	ipStartStrings := strings.Split(start, ".")
	ipFinalStrings := strings.Split(end, ".")

	fc := func(addr []string) []int {
		ipIntegers := make([]int, len(addr))

		for i, num := range addr {
			ipInt, _ := strconv.Atoi(num)
			ipIntegers[i] = ipInt
		}
		return ipIntegers
	}

	ipStartNums := fc(ipStartStrings)
	ipFinalNums := fc(ipFinalStrings)
	octetCount3 := int(math.Pow(256, 3))
	octetCount2 := int(math.Pow(256, 2))
	ipStartCount := (ipStartNums[0] * octetCount3) + (ipStartNums[1] * octetCount2) + (ipStartNums[2] * 256) + ipStartNums[3] + 1
	ipFinalCount := (ipFinalNums[0] * octetCount3) + (ipFinalNums[1] * octetCount2) + (ipFinalNums[2] * 256) + ipFinalNums[3] + 1

	return ipFinalCount - ipStartCount
}
