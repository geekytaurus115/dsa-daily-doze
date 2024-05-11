package main

import "fmt"

func main() {
	var (
		num int
	)

	fmt.Printf("Enter the number: ")
	//fmt.Scanln(&input)
	fmt.Scanf("%d", &num)

	factResult := 1
	for i := num; i >= 2; i-- {
		factResult *= i
	}

	fmt.Printf("Factorial of %d is %d\n", num, factResult)
	fmt.Printf("Factorial of %d is %d using Recurssion\n", num, factorial(num))

}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
