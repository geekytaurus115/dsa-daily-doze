package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		num     int
		isPrime bool
	)

	fmt.Print("Enter a number above 2: ")
	fmt.Scanf("%d", &num)

	isPrime = true

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Printf("The number %d is Prime", num)
	} else {
		fmt.Printf("The number %d is Not Prime", num)
	}
	fmt.Println()

}
