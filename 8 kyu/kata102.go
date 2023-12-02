package main

import "fmt"

func main() {
	var aNum, bNum, cNum int
	_, err := fmt.Scan(&aNum, &bNum, &cNum)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(ExpressionMatter(aNum, bNum, cNum))
}

func ExpressionMatter(a int, b int, c int) int {
	maxVariant := 0
	slice := make([]int, 6)
	slice[0] = a * (b + c)
	slice[1] = a * b * c
	slice[2] = a + b*c
	slice[3] = (a + b) * c
	slice[4] = a + b + c
	slice[5] = a*b + c

	for _, num := range slice {
		if num > maxVariant {
			maxVariant = num
		}
	}
	return maxVariant
}
