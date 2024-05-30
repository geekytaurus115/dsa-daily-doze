package main

import (
	"fmt"
	"sort"
)

func main() {

	var arrSize int
	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&arrSize)

	arr := make([]int, arrSize)

	fmt.Printf("Enter %d elements for the array: ", arrSize)
	for i := 0; i < arrSize; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)

	fmt.Println("Sorted array: ", arr)

}
