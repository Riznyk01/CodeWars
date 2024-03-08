package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Is_valid_ip("0.34.82.53"))
}
func Is_valid_ip(ip string) bool {
	ipArr := strings.Split(ip, ".")
	if len(ipArr) != 4 {
		return false
	}
	for _, o := range ipArr {
		n, err := strconv.Atoi(o)
		if err != nil {
			return false
		}
		if n > 255 {
			return false
		}
	}
	return true
}
