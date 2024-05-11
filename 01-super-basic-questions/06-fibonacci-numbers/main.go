package main

import "fmt"

func main() {
	var sequenceNumber int

	fmt.Print("Enter the sequence number to print fibonacci numbers: ")
	fmt.Scanf("%d", &sequenceNumber)

	sum0 := 0
	sum1 := 1
	result := 0

	if sequenceNumber == 0 {
		fmt.Print(sum0)
	} else if sequenceNumber == 1 {
		fmt.Print(sum1)
	} else {
		fmt.Print(sum0, ", ", sum1)
		for i := 3; i <= sequenceNumber; i++ {

			result = sum0 + sum1
			sum0 = sum1
			sum1 = result

			fmt.Printf(", %d", result)
		}
		fmt.Println()
	}
}
