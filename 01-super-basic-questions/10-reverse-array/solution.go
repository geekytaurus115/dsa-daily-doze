package main

import "fmt"

func reverseArray(arr []int) []int {

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]

	}

	return arr
}

func main() {

	//myArr := []int{1, 2, 3, 4, 5}

	var n int
	fmt.Println("Enter the size of your array: ")
	fmt.Scan(&n)

	fmt.Println("Enter the element of the array: ")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Reversed Array: ", reverseArray(arr))

}
