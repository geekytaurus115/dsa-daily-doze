package main

import "fmt"

func main() {
	var arrSize int

	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&arrSize)

	arr := make([]int, arrSize)

	fmt.Println("Enter the array elements: ")
	for i := 0; i < arrSize; i++ {
		fmt.Scan(&arr[i])
	}

	oddCount := 0
	evenCount := 0

	for i := 0; i < arrSize; i++ {
		if arr[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	fmt.Println("No. of odd elements: ", oddCount)
	fmt.Println("No. of even elements: ", evenCount)

}
