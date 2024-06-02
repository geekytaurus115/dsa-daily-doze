package main

import "fmt"

func main() {
	var arrSize int
	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&arrSize)

	arr := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		fmt.Scan(&arr[i])
	}

	sum := 0

	for i := 0; i < arrSize; i++ {
		sum += arr[i]
	}

	fmt.Println("Sum of arr elements: ", sum)

}
