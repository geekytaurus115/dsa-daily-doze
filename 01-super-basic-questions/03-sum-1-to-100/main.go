package main

import "fmt"

func main() {
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("Sum of all numbers between 1 and 100: ", sum)
}
