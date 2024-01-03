package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ModifyMultiply("nBkUGmB LzlHA xhrEyzEgzcDONLXZ boofJQGTlGb JDHLFsTfofdkPRdTuhoOTHar ESblTQGnSKhfYTuYiBL EaokFbOkMaJzVBlZLvDIM rxKsHtXwDBQUoDwIM wTr", 6, 0))
}
func ModifyMultiply(str string, loc, num int) string {
	if num == 0 {
		num++
	}
	arr := make([]string, num)
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.Split(str, " ")[loc]
	}
	return strings.Join(arr, "-")
}
