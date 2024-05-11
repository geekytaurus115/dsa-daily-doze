package main

import "fmt"

func main() {

	var num int
	fmt.Println("Enter a number: ")
	fmt.Scan(&num)

	original := num
	reverse := 0

	for num > 0 {
		digit := num % 10
		reverse = reverse*10 + digit
		num /= 10
	}

	if reverse == original {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Not Palindrom")
	}

}
