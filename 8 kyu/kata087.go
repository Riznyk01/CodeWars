package main

import (
	"fmt"
)

func main() {
	var aNum, bNum, cNum int
	_, err := fmt.Scan(&aNum, &bNum, &cNum)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(GetGrade(aNum, bNum, cNum))
}

func GetGrade(a, b, c int) rune {
	avg := (a + b + c) / 3
	switch {
	case avg >= 80 && avg < 90:
		return 'B'
	case avg >= 70 && avg < 80:
		return 'C'
	case avg >= 60 && avg < 70:
		return 'D'
	case avg >= 0 && avg < 60:
		return 'F'
	default:
		return 'A'
	}
}
