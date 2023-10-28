package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	fmt.Println(MultiTable(num))
}
func MultiTable(number int) string {
	var row string
	for i := 1; i <= 10; i++ {
		row += fmt.Sprintf("%d * %d = %d", i, number, i*number)
		if i < 10 {
			row += "\n"
		}
	}
	return row
}
