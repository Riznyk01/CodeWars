package main

import "fmt"

func main() {
	var num uint64

	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	fmt.Println(ProductFib(num))
}

func ProductFib(prod uint64) [3]uint64 {
	var n1, n2, sum uint64
	arr := [3]uint64{0, 0, 0}
	n1 = 1
	n2 = 1

	for i := 1; n1*n2 <= prod; i++ {
		sum = n1 + n2
		n1 = n2
		n2 = sum

		if n1*n2 == prod {
			arr[0] = n1
			arr[1] = n2
			arr[2] = 1
			break
		}
	}

	if arr[2] == 0 {
		arr[0] = n1
		arr[1] = n2
		arr[2] = 0
	}

	return arr
}
