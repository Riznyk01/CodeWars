package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the size of the array: ")

	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	if n < 1 {
		fmt.Println("The size of the array must be greater than 0")
		return
	}

	arr := make([]int, n)

	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		_, err := fmt.Scan(&arr[i])

		if err != nil {
			fmt.Println("Input error:", err)
			return
		}
	}
	fmt.Println("Outlier is:", FindOutlier(arr))
}

func FindOutlier(integers []int) int {
	var odd, even, out int
	for i := 0; even < 2 && odd < 2; i++ {
		if integers[i]%2 == 0 {
			even++
			if even == 2 {
				for i := 0; i < len(integers); i++ {
					if integers[i]%2 != 0 {
						out = integers[i]
						break
					}
				}
			}
		} else {
			odd++
			if odd == 2 {
				for i := 0; i < len(integers); i++ {
					if integers[i]%2 == 0 {
						out = integers[i]
						break
					}
				}
			}
		}
	}
	return out
}
