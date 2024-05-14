package main

import "fmt"

func main() {
	var num int

	fmt.Println("Enter the size of the Array: ")
	fmt.Scan(&num)

	arr := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&arr[i])

	}

	max1 := arr[0]
	max2 := arr[1]

	for _, num := range arr {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}

	}

	fmt.Println("Max-1:  ", max1)
	fmt.Println("Max-2:  ", max2)

}
